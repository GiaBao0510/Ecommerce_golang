package dto

type ConfirmEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type Email struct {
	From     string    `json:"from"`
	To       []Address `json:"to"`
	Subject  string    `json:"subject"`
	Text     string    `json:"text"`
	Category string    `json:"category"`
}

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}


