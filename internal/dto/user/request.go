package dto

type UserRequestByOAuth2 struct {
	Provider   string `json:"provider"`
	ProviderId string `json:"provider_id"`

	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	PhoneNum  string `json:"phone_num"`
	AvatarUrl string `json:"avatar_url"`
}

type PaginationRequest struct {
	Limit  int32 `form:"limit"  binding:"omitempty,min=1,max=100"`
	Offset int32 `form:"offset" binding:"omitempty,min=0"`
}

func (p *PaginationRequest) SetDefault() {
	if p.Limit == 0 {
		p.Limit = 20
	}
}
