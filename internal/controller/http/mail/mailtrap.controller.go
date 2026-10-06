package mail

import "github.com/gin-gonic/gin"

func (ctr *EmailController) SendEmail_Mailtrap(ctx *gin.Context) error {
	return ctr.emailSVC.SendEmail(ctx, nil)
}