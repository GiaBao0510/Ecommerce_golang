package dto

type ConfirmEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type EmailMessage struct {
	To       string
	Subject  string
	Text     string
	HTML     string
	Category string
}

type SendEmailRequest1 struct {
	To       string `json:"to" binding:"required,email"`
	Subject  string `json:"subject" binding:"required,max=200"`
	Text     string `json:"text" binding:"omitempty,max=2000"`
	HTML     string `json:"html" binding:"omitempty"`
	Category string `json:"category" binding:"omitempty,max=100"`
}

type SendNotificationRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required,max=200"`
	Message string `json:"message" binding:"required,max=2000"`
}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}
