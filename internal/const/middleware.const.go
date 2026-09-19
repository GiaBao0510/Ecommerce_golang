package _const

const (
	RealIPKey       = "real_ip"
	RequestIDKey    = "request_id"
	TraceIDKey      = "trace_id"
	X_Real_IP       = "X-Real-IP"
	X_Trace_ID      = "X-Trace-ID"
	X_Forwarded_For = "X-Forwarded-For"
)

// Hằng số liên quan đến phân quyền truy cập (Authorization)
const (
	RoleAdmin = 1
	RoleUser  = 2
)