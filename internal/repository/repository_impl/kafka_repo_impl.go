package repositoryimpl

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaEventPublisher_RepoImpl struct {
	writer *kafka.Writer
	log    *loghelper.DBLogger
}

// NewKafkaRepoImpl khởi tạo 1 Writer DÙNG CHUNG cho MỌI topic
// kafka-go cho phép không chỉ định topic cố định lúc tạo Writer
func NewKafkaRepoImpl(brokers []string, logger *loghelper.DBLogger) repository.IEventPublisher {
	writer := &kafka.Writer{
		Addr: kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne, // Chờ broker xác nhận đã ghi trước khi coi là publish thành công
	}

	return &KafkaEventPublisher_RepoImpl{writer: writer, log: logger}
}

func (p *KafkaEventPublisher_RepoImpl) Publish(ctx context.Context, topic, key string, message []byte) error {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key: []byte(key),
		Value: message,
	})
	if err != nil {
		p.log.LogError("Failed to publish message to Kafka", err,
			zap.String("topic", topic),
			zap.String("key", key),
		)
		return err
	}

	return nil
}

func (p *KafkaEventPublisher_RepoImpl) Close() error {
	return p.writer.Close()
}