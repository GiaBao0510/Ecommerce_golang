package email

import (
	"github.com/gin-gonic/gin"
)

type EmailControllerInterface interface {
	SendVerificationEmail(c *gin.Context) error
	VerifyEmail(c *gin.Context) error
	VerifyOTP_viaEmail(c *gin.Context) error
}
