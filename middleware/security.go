package middleware

// SIDE PROJECT ON PROGRESS

// import (
// 	"errors"
// 	"time"

// 	"github.com/crewdible/go-lib/encryption"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// )

// func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authHeader := c.Request().Header.Get("Authorization")

// 		if authHeader == "" {
// 			return errors.New("Authorization header is required")
// 		}

// 		// we parse our jwt token and check for validity against our secret
// 		token, err := encryption.VerifyToken(c)

// 		if err != nil {
// 			return errors.New("Error parsing token")
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)

// 		if !(ok && token.Valid) {
// 			return errors.New("Invalid token")
// 		}

// 		if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
// 			return errors.New("jwt is expired")
// 		}

// 		return next(c)
// 	}
// }
