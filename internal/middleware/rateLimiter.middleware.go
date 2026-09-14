package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limmiter *rate.Limiter // Limiter cho IP này
	lastSeen time.Time     // Thời điểm lần cuối client này gửi request
}

var (
	mu      sync.Mutex                 // mutex để đồng bộ hóa truy cập vào map clients
	clients = make(map[string]*Client) // map lưu trữ các limiter cho từng IP
)

// hàm này áp dụng rate limiting dựa trên IP của client. Mỗi IP sẽ có một limiter riêng, giới hạn số request trong một khoảng thời gian nhất định.
func getRateLimiter(key string, requestSec, burst int) *rate.Limiter {

	mu.Lock()         // Khóa mutex để đảm bảo an toàn khi truy cập map clients
	defer mu.Unlock() // Giải phóng mutex sau khi xong

	// Kiểm tra xem IP này đã có limiter chưa, nếu chưa thì tạo mới
	client, exists := clients[key]

	// Nếu tồn tại thì sẽ chỉ cập nhật thời điểm lần cuối client này gửi request và trả về limiter hiện tại
	if exists {
		// Cập nhật thời điểm lần cuối client này gửi request
		client.lastSeen = time.Now()
		return client.limmiter
	}

	// Nếu không tồn tại thì tạo mới
	limiter := rate.NewLimiter(
		rate.Limit(requestSec),
		burst,
	)

	clients[key] = &Client{
		limmiter: limiter,
		lastSeen: time.Now(),
	}
	return limiter
}

// hàm cleanUpClients này mục đích để dọn dẹp các limiter của các IP đã lâu không gửi request, tránh việc lưu trữ quá nhiều limiter không cần thiết
func CleanUpClients() {

	// Trong vòng 3 phút, nếu một IP không gửi request nào thì sẽ xóa limiter của IP đó khỏi map clients
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// Bộ lọc để ratelimit
func rateLimitMiddleware(requestSec, burst int, scope string) gin.HandlerFunc {
	return func(ctx *gin.Context){
		ip := ctx.GetString(_const.RealIPKey) // Lấy real_ip từ context (được set bởi RealIPMiddleware)
		if ip == "" {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid IP"})
			return
		}
		
		limiter := getRateLimiter(scope+":"+ip, requestSec, burst)
		
		// Kiểm tra xem client có thể gửi request hay không
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many requests. Please try again later.",
				"message": "Bạn đã gửi quá nhiều request trong một khoảng thời gian ngắn. Vui lòng thử lại sau.",
			})
		}

		ctx.Next()

	}
}

// Bộ lọc rate limiting middleware, áp dụng cho những route công khai (public routes).
func RateLimitingMiddlewareForPublicAccess() gin.HandlerFunc {
	return rateLimitMiddleware(
		global.Config.RateLimit.PerClient.Request_sec_public,
		global.Config.RateLimit.PerClient.Burst_public,
		"public",
	)
}

// Bộ lọc rate limiting middleware, áp dụng cho những route riêng tư (private routes).
func RateLimitingMiddlewareForPrivateAccess() gin.HandlerFunc {
	return rateLimitMiddleware(
		global.Config.RateLimit.PerClient.Request_sec_private,
		global.Config.RateLimit.PerClient.Burst_private,
		"private",
	)
}
