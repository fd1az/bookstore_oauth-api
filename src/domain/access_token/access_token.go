package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/fdiaz7/bookstore_oauth-api/src/utils/errors"
	"github.com/fdiaz7/bookstore_users-api/utils/crypto_utils"
)

const (
	expirationTime       = 24
	grantTypePassword    = "password"
	grantTypeCredentials = "credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	//Grant type = "password"
	Username string `json:"username"`
	Password string `json:"password"`
	//Grant type = "credentials"
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}

	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid User id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at AccessToken) IsExpired() bool {

	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

//WEB CLIENT - Client-Id : 123
//Android APP - Client-Id : 3214
