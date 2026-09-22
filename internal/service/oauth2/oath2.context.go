package oauth2

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
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

func (c *OAuth2ServiceContext) LoginWithOAuth2(
	ctx context.Context,
	req *models.CreateUsersRequestNonStrict,
) (*models.LoginResponse, error) {
	return c.oauth2ServiceStrategy.Login(ctx, req)
}