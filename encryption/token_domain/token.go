package token

type (
	AccessDetails struct {
		AccessUuid string `json:"access_uuid"`
		UserId     int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Role       string `json:"role"`
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
