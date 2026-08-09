package email

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

func (ctr *EmailController) SendVerificationEmail(c *gin.Context) error {

	// Đầu vào nhận email qua query parameter
	var Param dto.Email_Param
	if err := c.ShouldBindQuery(&Param); err != nil {
		return apperrors.NewBadRequestError("Invalid email parameter: " + err.Error())
	}

	// Gọi service để gửi mã đến email để xác thực
	if err := ctr.emailSVC.SendVerificationEmail(c, Param.Email); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Mã xác thực đã được gửi đến email của bạn", nil)
	return nil
}
