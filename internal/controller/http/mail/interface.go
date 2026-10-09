package mail

import "github.com/gin-gonic/gin"

type EmailControllerInterface interface {
	SendOTP(c *gin.Context) error
	SendNotification(c *gin.Context) error
}