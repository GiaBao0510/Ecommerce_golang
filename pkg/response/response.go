package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Mục đích chủ yếu của tệp tin này là để định nghĩa phản hồi chuẩn từ server đến client,

bao gồm mã lỗi và thông điệp tương ứng. Điều này giúp đảm bảo rằng tất cả
các phản hồi từ server đều có cấu trúc nhất quán và dễ dàng xử lý ở phía client.
*/
type SuccessResponse struct {
	Code    int         `json:"code"`           // Mã lỗi HTTP
	Message string      `json:"message"`        // Thông điệp chi tiết
	Data    interface{} `json:"data,omitempty"` // Dữ liệu trả về (nếu có)
}

type ErrorResponse struct {
	Code    int    `json:"code"`    // Mã lỗi HTTP
	Status  string `json:"status"`  // Thông điệp dành cho nhà phát triển
	Message string `json:"message"` // Thông điệp chi tiết
}

func Success_Response(c *gin.Context, code int, message string, data interface{}){
	c.JSON(code, SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Error_Response(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	})
}