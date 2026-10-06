package mail

import (
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/email"
)

type EmailController struct {
	emailSVC *service.IEmailProviderService
}

func NewEmailController(email_svc *service.IEmailProviderService) EmailControllerInterface {
	return &EmailController{
		emailSVC: email_svc,
	}
}