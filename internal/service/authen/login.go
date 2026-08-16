package authen

import "context"

type LoginUseCase struct {
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (l *LoginUseCase) Login(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
