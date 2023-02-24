package middleware

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func IdempotencyMiddleware(keyName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			idKey := r.URL.Query().Get("idKey")
			if idKey == "" {
				return next(c)
			}

			idKeyCookie, _ := r.Cookie(keyName)
			if idKeyCookie != nil && idKeyCookie.Value == idKey {
				resp, _ := c.Cookie("idempotentResp")
				if resp != nil {
					repsFinal, _ := base64.RawStdEncoding.DecodeString(resp.Value)
					c.Blob(http.StatusOK, "application/json", repsFinal)
				}
				return nil
			}

			c.Set("delayResponse", true)

			err := next(c)
			if err != nil {
				return err
			}

			cookieDuration := int(time.Minute.Seconds()) / 2

			c.SetCookie(&http.Cookie{
				Name:     keyName,
				Value:    idKey,
				HttpOnly: true,
				Secure:   true,
				MaxAge:   cookieDuration,
			})

			if resp, ok := c.Get("response").([]byte); ok {
				http.SetCookie(c.Response(), &http.Cookie{
					Name:     "idempotentResp",
					Value:    base64.RawStdEncoding.EncodeToString(resp),
					HttpOnly: true,
					Secure:   true,
					MaxAge:   cookieDuration,
				})
				return c.Blob(http.StatusOK, "application/json", resp)
			}

			return nil
		}
	}
}
