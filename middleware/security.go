package middleware

import (
	"net/http"

	"github.com/crewdible/go-lib/encryption"
	token "github.com/crewdible/go-lib/encryption/token_domain"
	_http "github.com/crewdible/go-lib/http"
	"github.com/labstack/echo/v4"
)

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

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var accessContext struct {
			Access *token.AccessDetails
			echo.Context
		}

		accessContext.Context = c
		accessContext.Access, err = encryption.ExtractTokenMetadata(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _http.MapBaseResponse("failed", err.Error(), nil))
		}
		return next(accessContext)
	}
}
