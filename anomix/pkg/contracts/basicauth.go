package contracts

import (
	"github.com/labstack/echo"
)

// BasicAuthCredentials defines credentials required for basic auth
//go:generate goscinny validator -t BasicAuthCredentials -f $GOFILE
type BasicAuthCredentials struct {
	UserName *string `required:"true"`
	Password *string `required:"true"`
}

// ExtractFromHTTP defines how to extract the data form request
func (ba *BasicAuthCredentials) ExtractFromHTTP(c echo.Context) *Error {
	request := c.Request()
	var ok bool
	username, password, ok := request.BasicAuth()
	if !ok {
		return ErrUnAuthorized("User name or password is not available")
	}
	ba.UserName, ba.Password = &username, &password
	return nil
}
