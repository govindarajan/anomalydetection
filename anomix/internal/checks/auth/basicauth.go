package auth

import (
	"github.com/govindarajan/anomalydetection/anomix/internal/checks"
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
)

type basicAuth struct {
	password string
	userID   string
}

func (ba basicAuth) UserID() string { return ba.userID }
func (ba basicAuth) Token() string  { return ba.password }

// Authorize checks if scopes are available for the given user
func (ba basicAuth) Authorize(ctx types.Context, scopes ...interface{}) (bool, *contracts.Error) {
	return true, nil
}

// GetUserInfo returns user information
func (ba basicAuth) GetUserInfo(ctx types.Context) checks.UserInfo {
	return ba
}

// BasicAuth authenticator function for http request
func BasicAuth(ctr *contracts.BasicAuthCredentials) Authenticater {
	return basicAuth{userID: *ctr.UserName, password: *ctr.Password}
}

// Authenticate checks if the user exists with the given credentials
func (ba basicAuth) Authenticate(ctx types.Context) (bool, *contracts.Error) {
	// TODO: Do the check.
	return true, nil
}

// Permissions will give permissions
func (ba basicAuth) Permissions(ctx types.Context, args ...interface{}) (map[string][]string, error) {
	return nil, nil
}
