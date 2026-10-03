package timing

import (
	"context"
	"sync"
	"time"
)

type ctxKey struct{}

type Collector struct {
	mu sync.Mutex						// để bảo vệ truy cập đồng thời vào bản ghi thời gian
	records map[string]time.Duration	// bản ghi thời gian cho các nhãn khác nhau
}

// NewCollector tạo ra một Collector mới và lưu trữ nó trong ngữ cảnh.
func NewCollector(ctx context.Context) (context.Context, *Collector){
	c := &Collector{records: make(map[string]time.Duration)}
	return  context.WithValue(ctx, ctxKey{}, c), c
}

// Hàm FronContext lấy Collector từ ngữ cảnh. Nếu không tìm thấy, nó trả về nil.
func FromContext(ctx context.Context) *Collector {
	c, _ := ctx.Value(ctxKey{}).(*Collector)
	return c
}

// Track bắt đầu bấm giờ và trả về hàm để dừng.
func Track(ctx context.Context, name string) func() {
	start := time.Now()

	return func(){
		if c := FromContext(ctx); c != nil {
			c.mu.Lock()
			c.records[name] += time.Since(start)	// cộng dồn nếu gọi nhiều lần
			c.mu.Unlock()
		}
	}
}

func(c *Collector) Snapshot() map[string]float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make(map[string]float64, len(c.records))
	for k, v := range c.records {
		out[k] = float64(v.Milliseconds())/ 1000.0	// chuyển đổi sang giây
	}
	return out
}