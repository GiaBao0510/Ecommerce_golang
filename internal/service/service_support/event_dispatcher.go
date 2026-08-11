package servicesupport

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

// ================ Event Dispatcher =====================
// Đây là một cơ chế để quản lý và phân phối các sự kiện trong ứng dụng. Nó cho phép các phần khác nhau
//  của ứng dụng đăng ký lắng nghe các sự kiện cụ thể và phản hồi khi những sự kiện đó xảy ra.
// Điều này giúp tách biệt các phần của ứng dụng, làm cho mã dễ bảo trì hơn và tăng khả năng mở rộng.
// Dùng Go channel và Worker pool để xử lý các sự kiện bất đồng bộ, giúp cải thiện hiệu suất và khả năng phản hồi của ứng dụng.
// ========================================================

// Event là cấu trúc dữ liệu đại diện cho một sự kiện trong hệ thống.
// Nó chứa thông tin về tên sự kiện, dữ liệu đi kèm, thời điểm xảy ra và Trace ID để theo dõi sự kiện xuyên suốt các dịch vụ.
type Event struct {
	Name      string    // Tên sự kiện, dùng để xác định loại sự kiện
	Payload   any       // Dữ liệu đi kèm với sự kiện, có thể là bất kỳ kiểu dữ liệu nào
	Timestamp time.Time // Thời điểm sự kiện xảy ra
	TraceID   string    // Trace ID để theo dõi sự kiện xuyên su
}

// EventHandler là một hàm xử lý sự kiện. Mỗi loại sự kiện sẽ có một handler tương ứng.
type EventHandler func(ctx context.Context, event Event) error

// Dispatcher là bộ phát sự kiện in-memory.
// Nó quản lý việc đăng ký các handler cho các sự kiện và phân phối các sự kiện đến các handler tương ứng.
type Dispatcher struct {
	eventChan   chan Event              // Kênh để gửi và nhận các sự kiện
	handlers    map[string]EventHandler // Bản đồ ánh xạ tên sự kiện đến handler tương ứng
	handlersMux sync.RWMutex            // Mutex để bảo vệ truy cập đồng thời vào bản đồ handlers
	workerCount int                     // Số lượng worker goroutine để xử lý sự kiện
	quit        chan struct{}           // Kênh để báo dừng dispatcher
	wg          sync.WaitGroup          // WaitGroup để chờ các worker kết thúc khi shutdown
}

func NewDispatcher(bufferSize, workerCount int) *Dispatcher {
	if bufferSize <= 0 {
		bufferSize = 1000
	}
	if workerCount <= 0 {
		workerCount = 5
	}

	return &Dispatcher{
		eventChan:   make(chan Event, bufferSize),
		handlers:    make(map[string]EventHandler),
		workerCount: workerCount,
		quit:        make(chan struct{}),
	}
}

// Register gắn handler cho một loại event cụ thể
//
// Ví dụ:
//
//	dispatcher.Register("order.created", func(ctx, event) error {
//	    // gửi email xác nhận
//	    return nil
//	})
func (d *Dispatcher) Register(eventName string, handler EventHandler) {
	d.handlersMux.Lock()         // khóa mutex để đảm bảo an toàn khi ghi vào bản đồ handlers
	defer d.handlersMux.Unlock() // giải phóng mutex sau khi hoàn thành thao tác ghi
	d.handlers[eventName] = handler
}

// Emit phát một sự kiện vào hàng đợi. Các worker sẽ lấy sự kiện từ hàng đợi và gọi handler tương ứng.
func(d *Dispatcher) Emit(event Event){
	event.Timestamp = time.Now() // Gán thời điểm hiện tại cho sự kiện

	// Sử dụng Select với default để tránh block nếu kênh eventChan đầy. Nếu kênh đầy, sự kiện sẽ bị bỏ qua.
	select {
	case d.eventChan <- event:
		// Sự kiện được gửi thành công vào kênh
	default:
		zap.L().Warn("Sự kiện bị bỏ qua vì kênh eventChan đầy",
		zap.String("event_name", event.Name), 
		zap.Any("payload", event.Payload),
		zap.String("trace_id", event.TraceID))
	}
}

// Start khởi động các worker để xử lý sự kiện từ kênh eventChan.
func (d *Dispatcher) Start() {
	for i := 0; i < d.workerCount; i++ {
		d.wg.Add(1)
		go d.worker(i)
	}
	zap.L().Info("Event dispatcher started", zap.Int("worker_count", d.workerCount))
}

// worker là vòng đời của một goroutine xử lý sự kiện. Nó liên tục lắng nghe kênh eventChan và gọi handler tương ứng cho mỗi sự kiện nhận được.
func (d *Dispatcher) worker(id int) {
	defer d.wg.Done() // Đánh dấu worker đã hoàn thành khi kết thúc

	for{ 
		select {
		case event := <-d.eventChan:
			d.processEvent(id, event) // Xử lý sự kiện nhận được
		case <-d.quit:
			zap.L().Info("Worker is shutting down", zap.Int("worker_id", id))
			return // Thoát vòng lặp để kết thúc worker	
		}
	}
}

// ProcessEvent xử lý một sự kiện cụ thể bằng cách tìm handler tương ứng và gọi nó.
func(d *Dispatcher) processEvent(workerID int, event Event){
	d.handlersMux.RLock() // Khóa mutex để đảm bảo an toàn khi đọc từ bản đồ handlers
	handler, exists := d.handlers[event.Name]
	d.handlersMux.RUnlock() // Giải phóng mutex sau khi đọc xong

	if !exists {
		zap.L().Warn("No handler registered for event", zap.String("event_name", event.Name), zap.String("trace_id", event.TraceID))
		return
	}

	// Recover để 1 event lỗi không làm sập worker. Nếu handler panic, recover sẽ bắt lỗi và log ra.
	defer func(){
		if r := recover(); r != nil {
			zap.L().Error("Panic in event handler",
				zap.Int("worker_id", workerID),
				zap.String("event_name", event.Name),
				zap.Any("panic", r),
				zap.String("trace_id", event.TraceID),
			)
		}
	}()

	// Tạo context với timeout để tránh handler chạy mãi. Nếu handler không hoàn thành trong thời gian quy định, nó sẽ bị hủy.
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel() // Giải phóng context sau khi xử lý xong

	if err := handler(ctx, event); err != nil {
		zap.L().Error("Error in event handler",
			zap.Int("worker_id", workerID),
			zap.String("event_name", event.Name),
			zap.Error(err),
			zap.String("trace_id", event.TraceID),
		)
	}
}

// Stop dừng dispatcher một cách "graceful" (chờ xử lý hết event trong channel)
func (d *Dispatcher) Stop() {
	close(d.quit) // Gửi tín hiệu dừng cho các worker
	d.wg.Wait()   // Chờ tất cả các worker kết thúc
	close(d.eventChan) // Đóng kênh eventChan sau khi tất cả worker đã dừng
	zap.L().Info("event dispatcher stopped gracefully")
}