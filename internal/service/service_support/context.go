package servicesupport

import (
	"context"
	"errors"
)

// ============================================
// CONTEXT KEYS - Định nghĩa key riêng để tránh xung đột với thư viện khác
// Dùng kiểu unexported type để đảm bảo chỉ package này mới có thể tạo key
// ============================================
type contextKey string

const (
	actorKey contextKey = "actor"
	traceIDKey contextKey = "trace_id"
)

// Actor chứa thông tin người dùng đang thực hiện request.
// Được inject vào context ở tầng Middleware sau khi verify JWT.
type Actor struct {
	Uuid string
	Email string
	RoleName string // "admin", "customer", "seller"
	IsGuest bool // true nếu request chưa login
	IPAddress string
}

// InjectActor: nhét actor vào context (được gọi ở Middleware/Controller)
func InjectActor(ctx context.Context, actor Actor) context.Context {
	return context.WithValue(ctx, actorKey, actor)
}

// InjectTraceID: nhét trade ID vào context (từ middleware tracing)
func InjectTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// GetActor lấy Actor từ context một cách AN TOÀN.
//
// QUAN TRỌNG: Không dùng ctx.Value().(Actor) vì có thể panic nếu context rỗng.
// Hàm này dùng "comma-ok idiom" (if actor, ok := ...) để đảm bảo an toàn.
func GetActor(ctx context.Context) (Actor, error) {
	if ctx == nil {
		return Actor{}, errors.New("conte")
	}

	actor, ok := ctx.Value(actorKey).(Actor)
	if !ok{
		// Trả về Guest actor thay vì error để service không bị gián đoạn
		return Actor{IsGuest: true}, nil
	}
	return actor, nil
}

// MustGetActor giống GetActor nhưng trả về error rõ ràng nếu bắt buộc phải có user login.
// Dùng cho các endpoint protected (ví dụ: CreateOrder, UpdateProfile)
func MustGetActor(ctx context.Context) (Actor, error) {
	actor, err := GetActor(ctx)
	if err != nil {
		return Actor{}, err
	}
	if actor.IsGuest {
		return Actor{}, errors.New("authentication required")
	}
	return actor, nil
}

// GetTraceID lấy trace ID để log (hữu ích khi debug với Loki/Grafana)
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if traceID, ok := ctx.Value(traceIDKey).(string); ok {
		return traceID
	}
	return ""
}