package encryption

import (
	"encoding/json"
	"fmt"

	// _redis "microservice/shared/pkg/database/redis"
	"os"
	"strings"
	"time"

	_token "github.com/crewdible/go-lib/encryption/token_domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateToken(userid int) (*_token.TokenDetails, error) {
	var err error
	td := &_token.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	accUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	td.AccessUuid = accUUID.String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = fmt.Sprintf("%s++%d", td.AccessUuid, userid)

	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// func CreateAuth(userid string, td *_token.TokenDetails) error {
// client := _redis.RedisManager()
// at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
// rt := time.Unix(td.RtExpires, 0)
// now := time.Now()

// errAccess := client.Set(client.Context(), td.AccessUuid, userid, at.Sub(now)).Err()
// if errAccess != nil {
// 	return errAccess
// }

// errRefresh := client.Set(client.Context(), td.RefreshUuid, userid, rt.Sub(now)).Err()
// if errRefresh != nil {
// 	return errRefresh
// }
// return nil
// }

// get the token from the request body
func ExtractToken(c echo.Context) string {
	bearToken := c.Request().Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c echo.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c echo.Context) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(c echo.Context) (*_token.AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var res *_token.AccessDetails
	err = json.Unmarshal(claimsJSON, &res)
	return res, err
}

// func FetchAuth(authD *_token.AccessDetails) (string, error) {
// client := _redis.RedisManager()
// userid, err := client.Get(client.Context(), authD.AccessUuid).Result()
// if err != nil {
// 	return "", err
// }

// userid, err = strconv.Atoi(userid)
// if err != nil {
// 	return "", err
// }

// if authD.UserId != userid {
// 	return "", errors.New("unauthorized")
// }
// return userid, nil
// }

// func DeleteAuth(givenUuid string) (int64, error) {
// 	client := _redis.RedisManager()
// 	deleted, err := client.Del(client.Context(), givenUuid).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }

// func DeleteTokens(authD *_token.AccessDetails) error {
// 	client := _redis.RedisManager()
// 	//get the refresh uuid
// 	refreshUuid := fmt.Sprintf("%s++%s", authD.AccessUuid, authD.UserId)
// 	//delete access token
// 	deletedAt, err := client.Del(client.Context(), authD.AccessUuid).Result()
// 	if err != nil {
// 		return err
// 	}
// 	//delete refresh token
// 	deletedRt, err := client.Del(client.Context(), refreshUuid).Result()
// 	if err != nil {
// 		return err
// 	}
// 	//When the record is deleted, the return value is 1
// 	if deletedAt != 1 || deletedRt != 1 {
// 		return errors.New("something went wrong")
// 	}
// 	return nil
// }
