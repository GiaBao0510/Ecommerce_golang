package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

// Khai báo biến toàn cục
var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092" // Địa chỉ của Kafka broker, có thể là một chuỗi các broker phân tách bằng dấu phẩy nếu có nhiều broker
	kafkaTopic = "user_topic_vip"  // Topic mặc định mà producer sẽ gửi dữ liệu đến và consumer sẽ đọc dữ liệu từ đó
)

/*
	Hàm này để cho Producer, nó sẽ tạo một Kafka writer để gửi dữ liệu đến một topic cụ thể.

Chúng ta có thể sử dụng hàm này trong các phần khác của ứng dụng để kết nối
*/
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL), // Địa chỉ của Kafka broker
		Topic:    topic,               // Tên topic mà producer sẽ gửi dữ liệu đến
		Balancer: &kafka.LeastBytes{}, // Cơ chế cân bằng tải, ở đây sử dụng LeastBytes để gửi dữ liệu đến partition có ít dữ liệu nhất
	}
}

/*
	Hàm này để cho Consumer, nó sẽ tạo một Kafka reader để đọc dữ liệu từ một topic cụ thể.

Chúng ta có thể sử dụng hàm này trong các phần khác của ứng dụng để kết nối
và tiêu thụ dữ liệu từ Kafka.
*/
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,           // Danh sách các broker Kafka. Ví dụ ("localhost:9092,localhost:9093")
		Topic:          topic,             // Tên topic mà consumer sẽ đọc
		GroupID:        groupID,           // ID nhóm consumer để quản lý offset và phân phối tải
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // Đây là thời gian giữa các lệnh commit offset sau khi đọc dữ liệu, ở đây là 1 giây
		StartOffset:    kafka.FirstOffset, // Đọc giá trị Offset ban đầu khi consumer bắt đầu lắng nghe
	})
}

// ---- Ví dụ tạo một Func thực hiện mua bán chứng khoán	----
type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"` // "buy" hoặc "sell"
}

// Hàm khởi tạo
func NewStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

// hàm thực hiện
func actionStock(c *gin.Context) {

	// Lấy thông tin từ request
	s := NewStock(c.Query("message"), c.Query("type"))

	// Chuyển đổi thông tin thành JSON
	body := make(map[string]interface{})
	body["message"] = s.Message
	body["type"] = s.Type

	jsonBody, _ := json.Marshal(body)

	// Tạo một Message Kafka với key là "action" và value là JSON đã được chuyển đổi
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// Viết message bới Producer vào Kafka topic
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Action sent to Kafka successfully",
	})
}

// Consumer hocngs mua ATC
func RegisterConsumerATC(id int) {
	kafkaGroupID := "consumer-group-"

	// Tại đây consumer đăng ký đọc dữ liệu từ topic nào đó và cũng tại groupID nào đó
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupID)
	defer reader.Close() // Giải phóng tài nguyên khi không còn sử dụng

	//Viết vòng lặp để hiên thị những người dùng nào hóng mua ATC
	fmt.Printf("\nConsumer(%d) is listening for messages...\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error reading message: %v\n", id, err)
		}
		fmt.Printf("Consumer(%d), hóng topic: %v, partition: %d, time: %d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

// hàm main
func main() {
	r := gin.Default()

	// Khởi tạo Kafka producer
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close() // Giải phóng tài nguyên khi không còn sử dụng

	// Hàm này sẽ đẩy action từ hàm actionStock, để đăng lý một messga trong kafka
	r.POST("action/stock", actionStock)

	// Đăng ký 2 udser để mua Stock trong ATC
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}
