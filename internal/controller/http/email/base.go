package email

import (
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/usecase/user_usercase"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type EmailController struct {
	emailSVC service.VerifyUserUsecase
	logger   *loghelper.DBLogger
}

func NewEmailController(
	emailSVC service.VerifyUserUsecase,
	logger   *loghelper.DBLogger,
) EmailControllerInterface {
	return &EmailController{
		emailSVC: emailSVC,
		logger:   logger,
	}
}