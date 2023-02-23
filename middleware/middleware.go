package middleware

import (
	"io"
	"net/http"
	"time"

	"github.com/crewdible/go-lib/errors"
	_http "github.com/crewdible/go-lib/http"
	"github.com/crewdible/go-lib/logs"
	"github.com/crewdible/go-lib/stringlib"
	"github.com/labstack/echo/v4"
)

type requestLog struct {
	Timestamp   string `json:"timestamp"`
	RequestID   string `json:"requestId"`
	URL         string `json:"url"`
	ContentType string `json:"contentType,omitempty"`
	Body        []byte `json:"body,omitempty"`
}

func ErrorAndLoggingHandler(serviceName string) func(functionName string) echo.MiddlewareFunc {
	return func(functionName string) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				logger := logs.NewLogger("api", serviceName, functionName)
				c.Set("logger", logger)
				defer func() {
					go logger.Flush()
				}()

				timeloc, _ := time.LoadLocation("Asia/Jakarta")
				now := time.Now().In(timeloc)

				var reqBody []byte
				req := c.Request()
				reqID := stringlib.GenerateRandString(5)
				contentType := req.Header.Get("Content-Type")
				url := req.Host + req.URL.RequestURI()

				reqBody, err := io.ReadAll(req.Body)
				if err != nil {
					return _http.RespondErrorJSON(c, errors.Errors(err, &errors.ErrorOption{
						HTTPCode: http.StatusBadRequest,
					}))
				}

				logger.Log("request", requestLog{
					RequestID:   reqID,
					URL:         url,
					ContentType: contentType,
					Body:        reqBody,
					Timestamp:   now.Format("2006-01-02 15:04:05"),
				})

				err = next(c)
				if err != nil {
					return _http.RespondErrorJSON(c, err)
				}

				if resp, ok := c.Get("response").([]byte); ok {
					logger.Log("respponse", resp)
				}

				return nil
			}
		}
	}
}
