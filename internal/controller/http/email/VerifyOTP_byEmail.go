package email

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

func (ctr *EmailController) VerifyOTP_viaEmail(c *gin.Context) error {

	// Đầu vào nhận email và OTP qua body
	input := dto.ConfirmEmailRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid request body: " + err.Error())
	}

	// Gọi service để xác thực email
	if err := ctr.emailSVC.VerifyOTP_viaEmail(c, input.Email, input.OTP); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "OTP verified successfully", nil)
	return nil
}