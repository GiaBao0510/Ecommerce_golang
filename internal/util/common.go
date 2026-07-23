package util

import (
	"fmt"
	"strconv"
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

// Hàm kiểm tra và chuyển đổi ID từ string sang int32
func VerifyID(id string) (int32, error) {
	id_int, err := strconv.Atoi(id)

	// Kiểm tra nếu có lỗi khi chuyển đổi ID từ string sang int
	if err != nil {
		return 0, apperrors.NewBadRequestError("ID phải là một số nguyên hợp lệ")
	}

	if id_int <= 0 {
		return 0, apperrors.NewBadRequestError("Mã ID phải lớn hơn 0")
	}

	return int32(id_int), nil
}

// Hàm xác minh tên không được để trống
func VerifyName(name string) error {
	if name == "" {
		return apperrors.NewBadRequestError("Status name cannot be empty")
	}
	return nil
}

// Hàm helpers để lấy string từ context, nếu không có thì trả về giá trị mặc định
func GetSringFromContextVlue(ctx *gin.Context, key string) string {
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
