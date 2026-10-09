package mail

import (
	"net/http"

	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/email"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

type EmailController struct {
	emailSVC service.IEmailService
}

func NewEmailController(email_svc service.IEmailService) EmailControllerInterface {
	return &EmailController{
		emailSVC: email_svc,
	}
}

// Gửi thông tin xác thực email đến người dùng - POST
func(ctr *EmailController) SendOTP(c *gin.Context) error {
	
	// lấy thông tin đầu vào
	var input dto.SendOTPRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return controller.HandleValidationError(err)
	}
	
	// Gọi service để gửi email xác thực
	if err := ctr.emailSVC.SendOTP(c, input.Email); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Verification email sent successfully", nil)
	return nil
}

// Gửi thông báo đến người dùng - POST
func (ctr *EmailController) SendNotification(c *gin.Context) error {
	// lất email trên query param
	var email dto.SendNotificationRequest
	if err := c.ShouldBindJSON(&email); err != nil {
		return controller.HandleValidationError(err)
	}

	// Gọi service để gửi email thông báo
	if err := ctr.emailSVC.SendNotification(c, email.Email, email.Message); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Email notification sent successfully", nil)
	return nil
}