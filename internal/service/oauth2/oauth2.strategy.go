package oauth2

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuth2UserInfo struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	AvatarURL     string `json:"avatar_url"`
	EmailVerified bool   `json:"email_verified"`
}

type IOAuth2ServiceStrategy interface {
	GetAuthURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfor(ctx context.Context, token *oauth2.Token) (*OAuth2UserInfo, error)
	//Login(ctx context.Context, req *models.CreateUsersRequestNonStrict) (*models.LoginResponse, error)
}