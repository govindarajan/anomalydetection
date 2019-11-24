package contracts

import (
	"strings"

	"github.com/labstack/echo"
)

// BearerAuthCredentials defines credentials required for basic auth
//go:generate goscinny validator -t BearerAuthCredentials -f $GOFILE
type BearerAuthCredentials struct {
	Token      *string `required:"true"`
	AccountSid *string
}

// ExtractFromHTTP defines how to extract the data form request
func (ba *BearerAuthCredentials) ExtractFromHTTP(c echo.Context) *Error {
	authHeaders := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(authHeaders) < 2 || strings.ToLower(authHeaders[0]) != "bearer" {
		return ErrUnAuthorized("bearer token is not found")
	}
	ba.Token = &authHeaders[1]
	if accountSid := c.Param("accountSid"); accountSid != "" {
		ba.AccountSid = &accountSid
	}
	return nil
}
