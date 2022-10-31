package token

type (
	AccessDetails struct {
		AccessUuid string
		UserId     int
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
