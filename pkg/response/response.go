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
type ResponseData struct {
	Code    int         `json:code`
	Message string      `json:message`
	Data    interface{} `json:data`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
