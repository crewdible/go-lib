package jwt

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/v4"
// )

// type TokenCrew struct {
// 	Authorized bool
// 	UserID     string
// 	Exp        int64
// 	Signature  string
// }

// type TokenResponse struct {
// 	AccessToken  string
// 	AtExpireIn   int
// 	RefreshToken string
// 	RtExpireIn   int
// }

// func getJwtSecret() string {
// 	if os.Getenv("JWT_SECRET") != "" {
// 		return os.Getenv("JWT_SECRET")
// 	} else {
// 		return "etiY8Dh-9ZwSUGblW5fyfuZ17vN8qFnd.QEG4mNsoW1PYtYSa0!Ud"
// 	}
// }

// func getJwtSignatureSecret() string {
// 	if os.Getenv("JWT_SIGNATURE_SECRET") != "" {
// 		return os.Getenv("JWT_SECRET")
// 	} else {
// 		return "etiY8Dh-9ZwSUGblKbda9wHIjH2ncYlFnd.QEG4mNsoPYtYSa0!Ud"
// 	}
// }

// func GenerateTokenResponse(userID string) *TokenResponse {

// }

// func TokenRefresh(c echo.Context) (*TokenResponse, error) {
// 	token, err := VerifyToken(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, err
// 	}

// 	claimsJSON, err := json.Marshal(claims)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var res *TokenCrew
// 	err = json.Unmarshal(claimsJSON, &res)

// 	at := CreateToken(res.UserID, res.Signature)
// 	return &TokenResponse{}, err
// }

// // abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
// func generateRefreshToken(created, expired string) string {
// 	var (
// 		n           = 10
// 		letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 		ts          = strconv.Itoa(int(time.Now().Unix()))
// 		es          = strconv.Itoa(int(time.Now().Add(time.Hour * 6).Unix()))
// 	)
// 	if created != "" {
// 		ts = created
// 	}

// 	if expired != "" {
// 		es = expired
// 	}

// 	replacer := strings.NewReplacer("0", "x", "1", "a", "2", "Z", "3", "d", "4", "K", "5", "t", "6", "S", "7", "f", "8", "Y", "9", "l")
// 	ts = replacer.Replace(ts)
// 	es = replacer.Replace(es)
// 	rt := fmt.Sprintf("%s-%s-", es, ts)
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	for i := 0; i < n; i++ {
// 		rt += string(letterBytes[rand.Int63()%int64(len(letterBytes))])
// 	}

// 	return rt
// }

// func CreateToken(userID, signature string) (string, error) {
// 	var err error
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["user_id"] = userID
// 	atClaims["signature"] = signature
// 	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	accessToken, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return accessToken, nil
// }

// func TokenHasAccess(c echo.Context, userID string) (bool, error) {
// 	data, err := ExtractTokenMetadata(c)
// 	if err != nil {
// 		return false, err
// 	}

// 	if data.Exp <= time.Now().Unix() || !data.Authorized || data.UserID != userID {
// 		return false, err
// 	}

// 	return true, err
// }

// func ExtractToken(c echo.Context) string {
// 	bearToken := c.Request().Header.Get("Authorization")
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }

// func VerifyToken(c echo.Context) (*jwt.Token, error) {
// 	tokenString := ExtractToken(c)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(getJwtSecret()), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return token, nil
// }

// func TokenValid(c echo.Context) error {
// 	token, err := VerifyToken(c)
// 	if err != nil {
// 		return err
// 	}
// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 		return err
// 	}
// 	return nil
// }

// func ExtractTokenMetadata(c echo.Context) (*TokenCrew, error) {
// 	token, err := VerifyToken(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok && !token.Valid {
// 		return nil, err
// 	}

// 	claimsJSON, err := json.Marshal(claims)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var res *TokenCrew
// 	err = json.Unmarshal(claimsJSON, &res)
// 	return res, err
// }
