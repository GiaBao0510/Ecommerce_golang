package middleware

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ValidateBody bind và validate request body theo struct tag,
// đặt SÁT TRƯỚC business handler — request không hợp lệ bị chặn
// Tại đây sẽ tạo 1 instance mới của cùng kiểu dữ liệu cho mỗi request, tránh race condition khi nhiều
func ValidationMiddleware(model interface{}) gin.HandlerFunc {
	
	// xác định kiểu dữ liệu gốc 1 lần lúc đăng ký (reflect chỉ đọc kiểu dữ liệu tĩnh không đụng đến giá trị)
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	
	return func(c *gin.Context) {
		
		// tạo 1 instance mới của cùng kiểu dữ liệu cho mỗi request, tránh race condition khi nhiều
		instance := reflect.New(modelType).Interface()
		
 		if err := c.ShouldBindJSON(instance); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   http.StatusBadRequest,
				"status": "Bad Request",
				"Message": "Invalid request body",
				"details": err.Error(),
			})
			return
		}

		c.Set("validated_body", instance)
		c.Next()
	}
}
