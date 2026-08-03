package util

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	Password string
	Cost     int
}

// hàm khởi tạo
func NewBcrypt(obj Bcrypt) *Bcrypt {

	// Mặc định đặt cost là 14 nếu không được cung cấp
	cost := 14

	// phạm vi cost hợp lệ từ 4 đến 31
	if obj.Cost < 4 || obj.Cost > 31 {
		obj.Cost = cost
	}

	return &Bcrypt{
		Password: obj.Password,
		Cost:     obj.Cost,
	}
}

// Hàm băm mật khẩu
func (b *Bcrypt) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(b.Password), b.Cost)

	return string(bytes), err
}

// Hàm so sánh mật khẩu
func (b *Bcrypt) CheckPasswordHash(Password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password))
	return err == nil
}