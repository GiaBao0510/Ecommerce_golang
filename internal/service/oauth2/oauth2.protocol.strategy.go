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

// Oauth2ProviderStrategy sẽ định nghĩa các phương thức mà các nhà cung cấp OAuth2 phải triển khai. Nó cho phép chúng ta có thể mở rộng và thêm các nhà cung cấp OAuth2 khác nhau mà không cần thay đổi mã nguồn chính.
type Oauth2ProviderStrategy interface {
	GetAuthURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfor(ctx context.Context, token *oauth2.Token) (*OAuth2UserInfo, error)
}

