package email

import (
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
)

type EmailController struct {
	emailSVC *service.VerifyUserUsecase
}

func NewEmailController(
	emailSVC *service.VerifyUserUsecase,
) EmailControllerInterface {
	return &EmailController{
		emailSVC: emailSVC,
	}
}