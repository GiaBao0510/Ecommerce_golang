package mail

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/gin-gonic/gin"
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/http"
)

func (ctr *EmailController) SendEmail_Mailjet(ctx *gin.Context) error {

	// Đọc thông tin đầu vào từ request body
	var email dto.Email
	if err := ctx.ShouldBindJSON(&email); err != nil {
		return controller.HandleValidationError(ctx)
	}

	return ctr.emailSVC.SendEmail(ctx, &email)
}