package util

import (
	"fmt"
	"math/big"
	mathrand "math/rand"
	"crypto/rand"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/gin-gonic/gin"
)

// Hàm định dạng thời gian uptime
func FotmatUptime(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}

// Hàm xác minh tên không được để trống
func VerifyName(name string) error {
	if name == "" {
		return apperrors.NewBadRequestError("Status name cannot be empty")
	}
	return nil
}

// Hàm helpers để lấy string từ context, nếu không có thì trả về giá trị mặc định
func GetStringFromContextValue(ctx *gin.Context, key string) string {
	value, exists := ctx.Get(key)
	if !exists {
		return ""
	}
	s, ok := value.(string)
	if !ok {
		return ""
	}
	return s
}

// Hàm tạo chuỗi số nguyên có độ dài là n ký tự
func GenerateRandomNumber(n int) string {
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = byte(mathrand.Intn(10)) + '0'
	}

	return string(result)
}

// Hàm tạo ra một chuỗi ngẫu nhiên có độ dài n ký tự
func GenerateRandom(n int) ([]byte, error) {

	const letters ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	b := make([]byte, n)
	for i := range b {

		// rand.Int đọc từ crypto/rand.Reader (entropy của OS) — không thể dự đoán
        idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return nil, err
		}

		b[i] = letters[idx.Int64()]
	}
	return b, nil
}
