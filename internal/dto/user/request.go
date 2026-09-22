package dto

type UserRequestByOAuth2 struct {
	Provider string `json:"provider"`
	ProviderId string `json:"provider_id"`
	
	Email string `json:"email"`
	UserName string `json:"user_name"`
	PhoneNum string `json:"phone_num"`
	AvatarUrl string `json:"avatar_url"`
}