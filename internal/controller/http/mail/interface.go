package mail

import "github.com/gin-gonic/gin"

type EmailControllerInterface interface {
	SendEmail_Mailtrap(c *gin.Context) error
	SendEmail_Mailjet(c *gin.Context) error
}