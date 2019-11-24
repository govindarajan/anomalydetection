package auth

type userDetails struct {
	userID    string
	authToken string
}

func (u userDetails) UserID() string    { return u.userID }
func (u userDetails) AuthToken() string { return u.authToken }
