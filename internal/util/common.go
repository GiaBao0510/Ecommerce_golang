package util

import (
	"math/rand/v2"
	"fmt"
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

// Hàm tạo chuỗi số nguyên có độ dài là 6 ký tự
func GenerateRandomNumber() string {
	var b [6]byte
	for i := range b {
		b[i] = byte(rand.IntN(10)) + '0'
	}

	return string(b[:])
}
