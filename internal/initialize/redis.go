package initialize

import (
	"context"
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background() // Conntext này giúp kiểm soát goroutine

func InitRedis() {

	// Thiết lập kết nối Redis sử dụng cấu hình từ global.Config.Redis
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", r.Address, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize, // Số lượng kết nối tối đa trong pool (Ví dụ nếu đặt là 10, thì sẽ có 10 connec kết nối trong mỗi cpu khả dụng )
	})

	// Kiểm tra kết nối bằng cách ping Redis. Nếu có lỗi, nó sẽ log lỗi và panic để dừng chương trình
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error.Error("InitRedis: Failed to connect to Redis", zap.Error(err))
		panic(fmt.Sprintf("InitRedis: Failed to connect to Redis: %v", err)) // Panic để dừng chương trình nếu có lỗi nghiêm trọng, đồng thời cung cấp thông tin chi tiết về lỗi
	}

	global.Redis = rdb
	global.Logger.Access.Info("InitRedis: Successfully connected to Redis")
}

func CloseRedis() {
	if global.Redis != nil {
		if err := global.Redis.Close(); err != nil {
			global.Logger.Error.Error("CloseRedis: Failed to close Redis connection", zap.Error(err))
		}
	}
}
