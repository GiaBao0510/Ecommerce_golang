package oauth2

import (
	"context"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IOAuth2ServiceStrategy interface {
	Login(ctx context.Context, req *models.CreateUsersRequestNonStrict) (*models.LoginResponse, error)
}