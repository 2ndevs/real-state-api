package presenters

type RefreshTokenPresenter struct{}

type RefreshTokenToHTTP struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (RefreshTokenPresenter) ToHTTP(token string, refreshToken string) RefreshTokenToHTTP {
	return RefreshTokenToHTTP{
		Token: token,
    RefreshToken: refreshToken,
	}
}
