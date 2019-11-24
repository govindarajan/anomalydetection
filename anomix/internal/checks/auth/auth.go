package auth

import (
	"strings"

	"github.com/govindarajan/anomalydetection/anomix/internal/checks"
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
	"github.com/labstack/echo"
)

const (
	authTypeBasic  = "basic"
	authTypeBearer = "bearer"
)

//Authenticater defines interface for auth
type Authenticater interface {
	Authenticate(types.Context) (bool, *contracts.Error)
	Authorize(types.Context, ...interface{}) (bool, *contracts.Error)
	GetUserInfo(types.Context) checks.UserInfo
	Permissions(types.Context, ...interface{}) (map[string][]string, error)
}

func GetAuthenticator(c echo.Context, supportedAuths map[string]struct{}) (Authenticater, *contracts.Error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil, contracts.ErrUnAuthorized()
	}
	authSplit := strings.Split(authHeader, " ")
	authType := strings.ToLower(authSplit[0])
	if _, ok := supportedAuths[authTypeBasic]; ok && authType == authTypeBasic {
		basicCreds := &contracts.BasicAuthCredentials{}
		if err := basicCreds.ExtractFromHTTP(c); err != nil {
			return nil, err
		}
		if err := basicCreds.Validate(); err != nil {
			return nil, err
		}
		return BasicAuth(basicCreds), nil
	}
	return nil, contracts.ErrUnAuthorized("unsupported authtype", authType)
}
