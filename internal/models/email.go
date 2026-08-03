package models

type EmailVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	Uuid  string `json:"uuid" binding:"required"`
}

type Email_PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type EmailTemplate struct {
	Subject string            `json:"subject"`
	Body    string            `json:"body"`
	Data    map[string]string `json:"data"`
}

type EmailData struct {
	To        string            `json:"to"`
	Name      string            `json:"name"`
	Subject   string            `json:"subject"`
	HTMLBody  string            `json:"html_body"`
	TextBody  string            `json:"text_body"`
	Variables map[string]string `json:"variables"`
}

//func template_Email