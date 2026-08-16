package authen

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
)

type RegisterController struct {
	svc    service.IAuthService
	logger *zap.Logger
}

func NewRegisterController(svc service.IAuthService, logger *zap.Logger) *RegisterController {
	return &RegisterController{svc: svc, logger: logger}
}

func (L *RegisterController) Register(ctx *gin.Context) error {
	
	input := models.CreateUsersRequest{}

	// Parse JSON body vào struct CreateUsersRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid request body: " + err.Error())
	}

	// Gọi service để thực hiện đăng ký người dùng
	
	if err := L.svc.Register(ctx, input); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, 200, "User registered successfully", nil)
	return nil
}
