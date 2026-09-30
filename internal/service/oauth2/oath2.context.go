package oauth2

import (
	"context"
	"golang.org/x/oauth2"
) 

type OAuth2ServiceContext struct {
	oauth2ServiceStrategy IOAuth2ServiceStrategy
}

func NewOAuth2ServiceContext(strategy IOAuth2ServiceStrategy) *OAuth2ServiceContext {
	return &OAuth2ServiceContext{
		oauth2ServiceStrategy: strategy,
	}
}

// Hàm SetStrategy cho phép thay đổi chiến lược OAuth2 tại runtime nếu cần thiết (ví dụ: chọn chiến lược dựa trên cấu hình hoặc yêu cầu của người dùng)
func (c *OAuth2ServiceContext) SetStrategy(strategy IOAuth2ServiceStrategy) {
	c.oauth2ServiceStrategy = strategy
}

// Hàm GetAuthURL để lấy URL xác thực từ chiến lược OAuth2 hiện tại. Nó nhận vào một tham số state để bảo vệ chống lại các cuộc tấn công CSRF.
func(c *OAuth2ServiceContext) GetAuthURL(state string) string {
	return c.oauth2ServiceStrategy.GetAuthURL(state)
}

// Hàm ExchangeCode để trao đổi mã xác thực lấy được từ OAuth2 provider để lấy access token. Nó nhận vào mã xác thực và trả về access token hoặc lỗi nếu có.
func(c *OAuth2ServiceContext) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.oauth2ServiceStrategy.Exchange(ctx, code)
}

// Hàm GetUserInfo để lấy thông tin người dùng từ OAuth2 provider bằng access token. Nó nhận vào access token và trả về thông tin người dùng hoặc lỗi nếu có.
func(c *OAuth2ServiceContext) GetUserInfo(ctx context.Context, token *oauth2.Token) (*OAuth2UserInfo, error) {
	return c.oauth2ServiceStrategy.GetUserInfor(ctx, token)
}

// func (c *OAuth2ServiceContext) LoginWithOAuth2(
// 	ctx context.Context,
// 	req *models.CreateUsersRequestNonStrict,
// ) (*models.LoginResponse, error) {
// 	return c.oauth2ServiceStrategy.Login(ctx, req)
// }