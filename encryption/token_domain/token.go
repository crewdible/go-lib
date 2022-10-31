package token

import "github.com/labstack/echo/v4"

type (
	AccessDetails struct {
		AccessUuid string `json:"access_uuid"`
		UserId     int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Role       string `json:"role"`
	}

	AccessDetailsEchoContext struct {
		echo.Context
		Access *AccessDetails
	}

	TokenDetails struct {
		AccessToken  string
		RefreshToken string
		AccessUuid   string
		RefreshUuid  string
		AtExpires    int64
		RtExpires    int64
	}
)
