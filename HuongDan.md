# Hướng dẫn tự sửa repository helpers và middleware cho project e-commerce_golang

Tài liệu này viết theo hướng "sinh viên tự làm được": mỗi mục có giải thích khái niệm, nhận xét code hiện tại, cách refactor từng bước và code mẫu.

Project hiện tại có cấu trúc liên quan như sau:

```text
internal/
  database/                 # code sqlc generate: database.Status, database.Role,...
  models/                   # model dùng trong tầng business/API: models.Status,...
  middleware/               # middleware Gin
  repository/
    repository_impl/        # implementation repository
  routers/                  # đăng ký route theo nhóm
  initialize/router.go      # nơi đang setup Gin engine chính
```

Lưu ý: đề bài nhắc `internal/router/router.go`, nhưng trong project hiện tại không thấy file đó. File tương đương cần sửa là `internal/initialize/router.go`. Ngoài ra có `internal/routers/routers.go`, nhưng file này giống một bản setup router khác và hiện không phải nơi ráp toàn bộ route chính.

---

## 1. DRY trong repository helpers

### 1.1. DRY là gì?

DRY là viết tắt của "Don't Repeat Yourself", nghĩa là "đừng lặp lại chính mình".

Trong code, vi phạm DRY thường xảy ra khi:

- Nhiều file có cùng một đoạn logic giống nhau.
- Một hàm ở file A bị file B gọi tạm vì thiếu chỗ dùng chung rõ ràng.
- Khi muốn đổi logic, bạn phải sửa nhiều nơi, dễ sót và sinh bug.

Với repository trong project này, các hàm như `toStatusModel`, `toRoleModel`, `toPermissionModel` có nhiệm vụ chuyển kiểu dữ liệu từ tầng database sang tầng model:

```go
database.Status -> models.Status
database.Role -> models.Role
database.Permission -> models.Permission
```

Đây là logic mapping. Mapping không xấu. Vấn đề nằm ở cách đặt nó.

### 1.2. Đánh giá code hiện tại có vi phạm DRY không?

Có dấu hiệu vi phạm DRY nhẹ đến vừa.

Hiện tại:

- `internal/repository/repository_impl/status_repo_impl.go` có helper `toStatusModel`.
- `internal/repository/repository_impl/roles_repo_impl.go` có helper `toRoleModel`.
- `internal/repository/repository_impl/permission_repo_impl.go` có helper `toPermissionModel`.
- `internal/repository/repository_impl/RolePermission_repo_impl.go` có helper `toRolePermissionModel`, nhưng toàn bộ file đang comment.
- `user_repo_impl.go` gần như rỗng.

Nếu mỗi helper chỉ được dùng đúng trong một file, việc để helper private trong file đó vẫn chấp nhận được. Nhưng nếu file repository khác gọi helper của file khác, hoặc sau này controller/service/test cũng cần chuyển DB sang model, thì nên tách ra package dùng chung.

Vấn đề thiết kế hiện tại:

- Helper nằm trong package `repositoryimpl`, không nói rõ đây là logic mapping dùng chung.
- Tên helper private như `toStatusModel` dễ bị trùng ý tưởng ở file khác.
- Nếu có repository khác muốn map `database.Status`, nó có thể gọi ké helper trong file status, làm code phụ thuộc lung tung.
- Có bug trong `toPermissionModel`: `Deleted_at` đang lấy `p.UpdatedAt`, đúng ra nên lấy `p.DeletedAt`.

Code hiện tại:

```go
Deleted_at:  p.UpdatedAt,
```

Nên sửa thành:

```go
Deleted_at:  p.DeletedAt,
```

### 1.3. Nên refactor theo hướng nào?

Nên tạo một package chuyên chuyển đổi dữ liệu, ví dụ:

```text
internal/mapper/
  status_mapper.go
  role_mapper.go
  permission_mapper.go
```

Package `mapper` có trách nhiệm rõ ràng:

- Nhận struct từ `internal/database`.
- Trả về struct từ `internal/models`.
- Không gọi database.
- Không ghi log.
- Không xử lý HTTP.
- Không xử lý business rule phức tạp.

Như vậy repository chỉ làm đúng việc của repository:

- Gọi query từ sqlc.
- Bắt lỗi database.
- Ghi log lỗi database.
- Gọi mapper để chuyển dữ liệu.
- Trả dữ liệu về service.

### 1.4. Bước 1: tạo package mapper

Tạo thư mục:

```text
internal/mapper
```

Tạo file `internal/mapper/status_mapper.go`:

```go
package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToStatusModel(s database.Status) models.Status {
	return models.Status{
		Id_status:   s.IDStatus,
		Name:        s.Name,
		Description: s.Description.String,
		Created_at:  s.CreatedAt,
		Updated_at:  s.UpdatedAt,
		Deleted_at:  s.DeletedAt,
	}
}
```

Tạo file `internal/mapper/role_mapper.go`:

```go
package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToRoleModel(r database.Role) models.Role {
	return models.Role{
		Role_id:     r.RoleID,
		Role_name:   r.RoleName,
		Description: r.Description.String,
	}
}
```

Tạo file `internal/mapper/permission_mapper.go`:

```go
package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToPermissionModel(p database.Permission) models.Permission {
	return models.Permission{
		Action_id:   p.ActionID,
		Action_name: p.ActionName,
		Description: p.Description.String,
		Created_at:  p.CreatedAt,
		Updated_at:  p.UpdatedAt,
		Deleted_at:  p.DeletedAt,
	}
}
```

Nếu muốn chuẩn bị luôn cho role-permission, tạo file `internal/mapper/role_permission_mapper.go`:

```go
package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToRolePermissionModel(rp database.RolePermission) models.Role_Permission {
	return models.Role_Permission{
		Role_id:       rp.RoleID,
		Permission_id: rp.ActionID,
		Created_at:    rp.CreatedAt,
		Updated_at:    rp.UpdatedAt,
		Deleted_at:    rp.DeletedAt,
	}
}
```

### 1.5. Bước 2: sửa repository để dùng mapper

Ví dụ với `status_repo_impl.go`.

Thêm import:

```go
import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
)
```

Sau đó xóa helper cũ:

```go
func toStatusModel(s database.Status) models.Status {
	// ...
}
```

Đổi chỗ dùng:

```go
result := toStatusModel(rows)
```

thành:

```go
result := mapper.ToStatusModel(rows)
```

Trong vòng lặp:

```go
for _, status := range query {
	statuses = append(statuses, mapper.ToStatusModel(status))
}
```

Tương tự với roles:

```go
result := mapper.ToRoleModel(rows)
```

```go
for _, role := range query {
	roles = append(roles, mapper.ToRoleModel(role))
}
```

Tương tự với permission:

```go
result := mapper.ToPermissionModel(row)
```

```go
for _, v := range query {
	permissions = append(permissions, mapper.ToPermissionModel(v))
}
```

### 1.6. Bước 3: chạy format và test build

Sau khi sửa code Go, chạy:

```powershell
gofmt -w internal/mapper internal/repository/repository_impl
go test ./...
```

Nếu `go test ./...` lỗi vì phần khác của project chưa hoàn chỉnh, tối thiểu chạy:

```powershell
go test ./internal/repository/...
go test ./internal/middleware/...
```

### 1.7. Khi nào không cần tách mapper?

Không phải cứ thấy helper là phải tách. Nếu helper:

- Chỉ dùng trong đúng một file.
- Không bị copy-paste ở nơi khác.
- Không có khả năng dùng lại.

thì để private trong file cũng ổn.

Nhưng với project e-commerce, các entity như status, role, permission, user, product sẽ xuất hiện nhiều. Tách mapper sớm giúp code rõ hơn và dễ mở rộng hơn.

---

## 2. Middleware trong `internal/middleware`

### 2.1. Middleware là gì?

Middleware là một hàm chạy trước hoặc sau handler chính.

Ví dụ request đi vào API:

```text
Client -> Middleware 1 -> Middleware 2 -> Controller -> Service -> Repository
```

Trong Gin, middleware thường có dạng:

```go
func SomeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// xử lý trước handler
		c.Next()
		// xử lý sau handler
	}
}
```

Nếu middleware muốn dừng request, dùng:

```go
c.Abort()
return
```

Nếu muốn trả lỗi rồi dừng:

```go
c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
c.Abort()
return
```

### 2.2. Pipeline chuẩn đề bài đưa ra

Pipeline chuẩn:

```text
Recovery
-> RequestID
-> RealIP
-> Logger
-> Metrics
-> Tracing
-> Timeout
-> RateLimiter
-> CORS
-> SecurityHeaders
-> Auth
-> Authorization
-> Validation
-> Business Handler
```

Giải thích nhanh:

- `Recovery`: bắt panic để server không crash.
- `RequestID`: tạo ID cho mỗi request, giúp tra log.
- `RealIP`: lấy IP thật khi app chạy sau proxy/load balancer.
- `Logger`: ghi log request/response.
- `Metrics`: đo số request, latency, status code.
- `Tracing`: theo dõi request đi qua nhiều service.
- `Timeout`: giới hạn thời gian xử lý request.
- `RateLimiter`: giới hạn số request từ client.
- `CORS`: cho phép frontend domain khác gọi API.
- `SecurityHeaders`: thêm header bảo mật.
- `Auth`: xác thực người dùng là ai.
- `Authorization`: kiểm tra người dùng được phép làm gì.
- `Validation`: kiểm tra input.
- `Business Handler`: controller xử lý nghiệp vụ.

### 2.3. Đánh giá middleware hiện có

Trong `internal/middleware` hiện có:

```text
authen.middleware.go
author.middleware.go
cors.middleware.go
errorHandler.middleware.go
hsts.middleware.go
https_redirection.middleware.go
logger.middleware.go
reatelimit.middleware.go
routing.middleware.go
traceid.middleware.go
```

Nhận xét từng file:

| File                              | Tình trạng              | Nhận xét                                                                                                                          |
| --------------------------------- | ----------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `traceid.middleware.go`           | Có logic                | Tương đương RequestID. Tên nên thống nhất thành `RequestIDMiddleware` hoặc giữ `TraceID_Middleware` nhưng hiểu đây là request id. |
| `logger.middleware.go`            | Có logic                | Đúng ý tưởng, chạy sau RequestID để log có trace id.                                                                              |
| `errorHandler.middleware.go`      | Có logic nhưng chưa tốt | Đang recover panic nhưng không trả response 500 và có thể trùng vai trò với `gin.Recovery()`.                                     |
| `authen.middleware.go`            | Có logic demo           | Đang check token cứng `"valid_token"`, chưa phù hợp production. Vì project chưa có đăng ký/đăng nhập nên chưa nên bật global.     |
| `author.middleware.go`            | Rỗng                    | Chưa có authorization.                                                                                                            |
| `cors.middleware.go`              | Rỗng                    | Chưa có CORS.                                                                                                                     |
| `reatelimit.middleware.go`        | Rỗng                    | Chưa có rate limiter. Tên file đang gõ sai `reatelimit`, nên đổi thành `rate_limit.middleware.go` nếu muốn sạch.                  |
| `hsts.middleware.go`              | Rỗng                    | Chưa có HSTS/security header.                                                                                                     |
| `https_redirection.middleware.go` | Rỗng                    | Chưa redirect HTTPS. Chỉ nên bật khi deploy có HTTPS rõ ràng.                                                                     |
| `routing.middleware.go`           | Rỗng                    | Chưa rõ mục đích, có thể xóa nếu không dùng.                                                                                      |

Thiếu so với pipeline chuẩn:

- RealIP middleware rõ ràng.
- Metrics.
- Tracing đúng nghĩa distributed tracing.
- Timeout.
- CORS implementation.
- Security headers đầy đủ.
- Validation middleware hoặc cách validation thống nhất.

Thừa hoặc chưa nên bật:

- `AuthenMiddleware` chưa nên bật global vì chưa có login/register/JWT thật.
- `ErrorHandlerMiddleware` và `gin.Recovery()` đang bị trùng ý tưởng recovery.
- `https_redirection` chưa nên bật ở local dev nếu chưa cấu hình HTTPS.
- `routing.middleware.go` chưa có mục đích rõ.

### 2.4. Vấn đề trong thứ tự đăng ký hiện tại

Trong `internal/initialize/router.go`, hiện thứ tự đang là:

```go
r.Use(
	middleware.TraceID_Middleware(),
	middleware.HttpLoggerMiddleware(global.Logger.Access),
	gin.Recovery(),
)
```

Thứ tự này chưa khớp pipeline chuẩn vì `Recovery` nên đứng đầu để bắt panic từ tất cả middleware phía sau.

Nên sửa thành:

```text
Recovery -> RequestID/TraceID -> RealIP -> Logger -> Timeout -> RateLimiter -> CORS -> SecurityHeaders -> Handler
```

Do project chưa có auth thật, chưa bật:

```text
Auth -> Authorization
```

Khi nào có login/register/JWT thật, thêm Auth sau SecurityHeaders và chỉ áp dụng cho route cần bảo vệ.

### 2.5. Bổ sung middleware cần thiết

#### 2.5.1. Recovery

Bạn có thể dùng `gin.CustomRecovery` thay vì vừa dùng `ErrorHandlerMiddleware` vừa dùng `gin.Recovery()`.

Tạo hoặc sửa `internal/middleware/recovery.middleware.go`:

```go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		logger.Error("panic recovered",
			zap.Any("panic", recovered),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	})
}
```

Sau đó không cần dùng `ErrorHandlerMiddleware()` nữa.

#### 2.5.2. RequestID / TraceID

File `traceid.middleware.go` đang dùng được. Nhưng nên đổi tên hàm cho dễ hiểu:

```go
func RequestIDMiddleware() gin.HandlerFunc {
	return TraceID_Middleware()
}
```

Hoặc đổi hẳn `TraceID_Middleware` thành `RequestIDMiddleware`.

Nếu muốn giữ code hiện tại, vẫn dùng:

```go
middleware.TraceID_Middleware()
```

#### 2.5.3. RealIP

Gin có `c.ClientIP()`, nhưng khi chạy sau reverse proxy, cần cấu hình trusted proxies ở router.

Trong `internal/initialize/router.go`, sau khi tạo `r := gin.New()`:

```go
_ = r.SetTrustedProxies([]string{"127.0.0.1"})
```

Khi deploy thật sau Nginx/Docker network, thay bằng IP hoặc CIDR của proxy.

Nếu muốn có middleware lưu IP vào context:

Tạo `internal/middleware/real_ip.middleware.go`:

```go
package middleware

import "github.com/gin-gonic/gin"

const RealIPKey = "real_ip"

func RealIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(RealIPKey, c.ClientIP())
		c.Next()
	}
}
```

#### 2.5.4. Logger

`HttpLoggerMiddleware` hiện dùng được. Nhưng nên chú ý:

- Đặt sau RequestID để có trace id.
- Đặt sau RealIP nếu muốn log IP đã chuẩn hóa.
- Nếu `ctx.FullPath()` rỗng với route không match, có thể dùng thêm `ctx.Request.URL.Path`.

Có thể chỉnh nhẹ:

```go
path := ctx.FullPath()
if path == "" {
	path = ctx.Request.URL.Path
}
```

#### 2.5.5. Timeout

Timeout giúp request không chạy mãi.

Tạo `internal/middleware/timeout.middleware.go`:

```go
package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"message": "Request timeout",
			})
			return
		}
	}
}
```

Ghi chú quan trọng: kiểu timeout dùng goroutine như trên dễ hiểu cho sinh viên, nhưng trong production cần cẩn thận vì handler vẫn có thể đang ghi response. Cách an toàn hơn là đặt timeout ở server HTTP hoặc dùng thư viện middleware đã được kiểm thử. Với bài học này, mục tiêu là hiểu khái niệm.

#### 2.5.6. RateLimiter

Rate limiter giới hạn số request để giảm spam/brute force.

Vì project đã có Redis trong `internal/initialize/redis.go`, sau này có thể làm Redis rate limit. Bản đơn giản cho sinh viên có thể dùng memory trước.

Tạo `internal/middleware/rate_limit.middleware.go`:

```go
package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	count     int
	resetTime time.Time
}

func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	var mu sync.Mutex
	visitors := make(map[string]*visitor)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		mu.Lock()
		v, exists := visitors[ip]
		if !exists || now.After(v.resetTime) {
			v = &visitor{
				count:     0,
				resetTime: now.Add(window),
			}
			visitors[ip] = v
		}

		v.count++
		allowed := v.count <= limit
		mu.Unlock()

		if !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
```

Nhược điểm của bản memory:

- Restart server là mất dữ liệu.
- Chạy nhiều instance thì mỗi instance đếm riêng.
- Map có thể lớn dần nếu nhiều IP lạ. Sau này nên thêm cleanup hoặc dùng Redis.

#### 2.5.7. CORS

CORS cho phép frontend khác origin gọi API. Ví dụ frontend chạy `http://localhost:3000`, backend chạy `http://localhost:8080`.

Sửa `internal/middleware/cors.middleware.go`:

```go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Trace-ID")
		c.Header("Access-Control-Expose-Headers", "X-Trace-ID")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
```

Không nên để `"*"` nếu dùng cookie hoặc credential.

#### 2.5.8. SecurityHeaders

Security headers giúp giảm một số rủi ro phổ biến trên trình duyệt.

Sửa hoặc tạo `internal/middleware/security_headers.middleware.go`:

```go
package middleware

import "github.com/gin-gonic/gin"

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "0")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	}
}
```

Về HSTS:

Chỉ bật HSTS khi website đã chạy HTTPS ổn định. Nếu bật HSTS lúc local/dev hoặc domain chưa HTTPS chuẩn, trình duyệt có thể ép HTTPS và làm bạn khó test.

Nếu cần HSTS, tạo:

```go
func HSTSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Next()
	}
}
```

#### 2.5.9. Auth và Authorization

Hiện `AuthenMiddleware` đang check:

```go
if token != "valid_token" {
```

Đây chỉ là demo, chưa phải authentication thật.

Vì project chưa có đăng ký/đăng nhập, không nên bật middleware này global. Nếu bật global, cả `/health`, `/checkStatus`, hoặc API public cũng bị chặn.

Nên làm như sau:

- Tạm thời không đăng ký `AuthenMiddleware` ở global `r.Use`.
- Khi có login/register/JWT, sửa middleware để parse Bearer token.
- Áp dụng auth theo group cần bảo vệ, ví dụ `/manager`.

Ví dụ sau này:

```go
ManagerGroup := MainGroup.Group("/manager")
ManagerGroup.Use(middleware.AuthMiddleware())
ManagerGroup.Use(middleware.AuthorizationMiddleware("manager"))
```

#### 2.5.10. Validation

Validation nên đặt gần handler vì mỗi endpoint có struct input khác nhau.

Ví dụ trong controller:

```go
var req models.Status
if err := c.ShouldBindJSON(&req); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid request body",
		"error":   err.Error(),
	})
	return
}
```

Nếu muốn tách middleware validation chung, có thể viết generic sau, nhưng với sinh viên nên làm trong controller trước cho dễ hiểu.

### 2.6. Thứ tự đăng ký middleware đúng trong `internal/initialize/router.go`

Sửa đoạn `r.Use(...)` trong `internal/initialize/router.go` theo hướng này:

```go
r.Use(
	// 1. Recovery: đứng đầu để bắt panic từ middleware/handler phía sau
	middleware.RecoveryMiddleware(global.Logger.Error),

	// 2. RequestID/TraceID: tạo id cho request
	middleware.TraceID_Middleware(),

	// 3. RealIP: lấy IP client
	middleware.RealIPMiddleware(),

	// 4. Logger: log request, cần trace id và IP
	middleware.HttpLoggerMiddleware(global.Logger.Access),

	// 5. Timeout: tránh request chạy quá lâu
	middleware.TimeoutMiddleware(5*time.Second),

	// 6. Rate limiter: hạn chế spam
	middleware.RateLimitMiddleware(100, time.Minute),

	// 7. CORS: xử lý preflight OPTIONS trước khi vào route
	middleware.CORSMiddleware(),

	// 8. Security headers: thêm header bảo mật
	middleware.SecurityHeadersMiddleware(),
)
```

Nhớ thêm import `time`:

```go
import (
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/middleware"
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers"
	"github.com/gin-gonic/gin"
)
```

Sau khi có auth thật, không nên thêm auth vào global ngay. Nên bảo vệ từng group:

```go
ManagerGroup := MainGroup.Group("/manager")
ManagerGroup.Use(middleware.AuthMiddleware())
ManagerGroup.Use(middleware.AuthorizationMiddleware("manager"))
```

Còn route public thì để ngoài auth:

```go
MainGroup.GET("/checkStatus", ...)
UserGroup := MainGroup.Group("/user")
```

### 2.7. Có cần sửa `internal/routers/routers.go` không?

File `internal/routers/routers.go` hiện cũng tạo `gin.New()` và đăng ký middleware:

```go
func SetUpRouter(logger *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(...)
}
```

Nhưng project chính đang dùng `internal/initialize/router.go` để setup router đầy đủ với manager/user routes.

Bạn nên chọn một nơi duy nhất để setup Gin engine:

- Cách đơn giản: giữ `internal/initialize/router.go` là nơi chính.
- Không đăng ký middleware global ở nhiều nơi.
- Nếu `SetUpRouter` không dùng nữa, có thể xóa sau khi chắc chắn không nơi nào gọi.

Kiểm tra nơi gọi:

```powershell
rg "SetUpRouter|InitRouter"
```

Nếu chỉ `InitRouter` được dùng trong `cmd/server/main.go`, thì tập trung sửa `InitRouter`.

---

## 3. Các vấn đề khác phát hiện trong project

### 3.1. Encoding comment tiếng Việt đang bị lỗi

Nhiều comment đang bị lỗi font, ví dụ:

```text
Triá»ƒn khai cÃ¡c phÆ°Æ¡ng thá»©c
```

Đây thường là lỗi encoding UTF-8 bị đọc sai. Code vẫn có thể chạy, nhưng rất khó học và khó bảo trì.

Cách xử lý:

- Dùng editor lưu file dưới dạng UTF-8.
- Sửa dần comment khi đụng vào file.
- Không cần sửa toàn bộ một lần nếu dễ gây diff lớn.

### 3.2. `ErrorHandlerMiddleware` chưa trả response cho client

Hiện middleware recover panic rồi log:

```go
if err := recover(); err != nil {
	global.Logger.Error.Error("Panic recovered...", ...)
}
```

Nhưng không trả JSON 500 rõ ràng. Client có thể nhận response không đúng mong muốn.

Nên thay bằng `RecoveryMiddleware` dùng `gin.CustomRecovery` như mục trên, rồi bỏ `ErrorHandlerMiddleware` khỏi global middleware.

### 3.3. `AuthenMiddleware` chưa phải auth thật

Hiện code:

```go
token := c.GetHeader("Authorization")
if token != "valid_token" {
```

Vấn đề:

- Token hard-code.
- Không hỗ trợ format `Bearer <token>`.
- Không parse JWT.
- Không set user id/role vào context.
- Nếu bật global sẽ chặn cả route public.

Tạm thời nên comment, chỉ bật khi có chức năng đăng ký/đăng nhập.

### 3.4. `Create` trong repository đang trả `0`

Ví dụ:

```go
func (s *StatusRepository) Create(ctx context.Context, obj *models.Status) (int, error) {
	// ...
	return 0, nil
}
```

Nếu interface yêu cầu trả ID bản ghi mới, thì query SQL nên dùng `RETURNING id_status`, sqlc generate hàm trả về ID, sau đó repository trả ID thật.

Ví dụ ý tưởng SQL:

```sql
INSERT INTO status (name, description)
VALUES ($1, $2)
RETURNING id_status;
```

Sau đó repository:

```go
id, err := s.db.CreateStatus(ctx, params)
if err != nil {
	return 0, apperrors.NewInternalServerError(err)
}
return int(id), nil
```

Nếu không cần trả ID, nên đổi interface thành `Create(ctx, obj) error` để code không gây hiểu nhầm.

### 3.5. Không nên trả NotFound khi danh sách rỗng trong mọi trường hợp

Hiện `GetAll` nếu không có dữ liệu thì trả lỗi NotFound:

```go
if len(query) == 0 {
	return nil, apperrors.NewNotFoundError(...)
}
```

Với API list, thông thường danh sách rỗng vẫn là thành công:

```json
[]
```

Nên cân nhắc đổi thành:

```go
if len(query) == 0 {
	return []models.Status{}, nil
}
```

Điều này giúp frontend dễ xử lý hơn: không có dữ liệu không nhất thiết là lỗi.

### 3.6. Tên file và tên hàm nên thống nhất

Một số tên hiện chưa thống nhất:

- `reatelimit.middleware.go` nên là `rate_limit.middleware.go`.
- `authen.middleware.go` nên là `auth.middleware.go`.
- `author.middleware.go` nên là `authorization.middleware.go`.
- `TraceID_Middleware` nên là `TraceIDMiddleware` hoặc `RequestIDMiddleware`.
- `Update_Put`, `Update_Patch` nên cân nhắc `UpdatePut`, `UpdatePatch` theo style Go.

Go thường dùng CamelCase, không dùng dấu `_` trong tên hàm exported.

### 3.7. File `RolePermission_repo_impl.go` đang comment gần hết

File này hiện có rất nhiều code comment. Nếu chưa làm, nên để TODO ngắn gọn thay vì comment cả block dài.

Ví dụ:

```go
package repositoryimpl

// TODO: implement RolePermissionRepository after SQL queries are ready.
```

Nếu để block comment lớn, người đọc khó biết code đó là bản cũ, bản nháp hay đang dùng.

### 3.8. Nên tránh setup router ở nhiều nơi

Hiện có:

- `internal/initialize/router.go`
- `internal/routers/routers.go`

Cả hai đều có dấu hiệu setup `gin.Engine`. Nên thống nhất:

- `initialize/router.go`: tạo Gin engine, đăng ký global middleware, gắn các route group.
- `routers/...`: chỉ nhận `*gin.RouterGroup` và đăng ký endpoint cụ thể.

Mô hình sạch:

```text
initialize/router.go
  tạo gin.Engine
  r.Use(global middleware)
  tạo /v1/api
  gọi routers.RouterGroupApp...

routers/manager/status.router.go
  chỉ đăng ký endpoint status
```

### 3.9. Thứ tự làm đề xuất

Nếu bạn muốn tự sửa theo thứ tự ít rủi ro, làm như sau:

1. Tạo package `internal/mapper`.
2. Chuyển `toStatusModel`, `toRoleModel`, `toPermissionModel` sang mapper.
3. Sửa bug `Deleted_at: p.DeletedAt`.
4. Chạy `gofmt`.
5. Tạo middleware còn thiếu: recovery, real_ip, timeout, rate_limit, cors, security_headers.
6. Sửa `internal/initialize/router.go` để đăng ký middleware đúng thứ tự.
7. Tạm thời không bật auth global.
8. Chạy `go test ./...`.
9. Nếu lỗi, sửa từng package theo log.

