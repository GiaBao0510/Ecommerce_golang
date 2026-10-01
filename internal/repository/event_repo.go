package repository

import "context"

// IEventPublisher trừu tượng hóa cho việc xuất bản sự kiện, cho phép các triển khai khác nhau (ví dụ: Kafka, RabbitMQ) được sử dụng mà không cần thay đổi mã gọi.	
type IEventPublisher interface {
	Publish(ctx context.Context, topic, key string, message []byte) error
	Close() error
}