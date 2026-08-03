package repositoryimpl

import (
	"context"
	"time"
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisRepositoryImpl struct {
	log *loghelper.DBLogger
}

// khởi tạo
func NewRedisRepositoryImpl(logger *loghelper.DBLogger) repository.IRedisRepository {
	return &RedisRepositoryImpl{log: logger}
}

func (r *RedisRepositoryImpl) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	err := global.Redis.Set(ctx, key, value, expiration).Err()
	if err != nil {
		r.log.LogError("Lôi khi set key-value vào Redis.", err, zap.String("key", key))
	}

	r.log.LogInfo("Set key-value", "Thực hiện thành công.",zap.String("key", key), zap.String("value", value), zap.Duration("expiration", expiration))
	return nil
}

func (r *RedisRepositoryImpl) Get(ctx context.Context, key string) (string, error) {
	value, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.log.LogWarning("Key không tồn tại trong Redis.", "Key không tồn tại.", zap.String("key", key))
			return "", nil // Trả về nil nếu key không tồn tại
		}
		r.log.LogError("Lỗi khi lấy giá trị từ Redis.", err, zap.String("key", key))
		return "", err
	}

	r.log.LogInfo("Get key-value", "Thực hiện thành công.", zap.String("key", key), zap.String("value", value))
	return value, nil
}

func (r *RedisRepositoryImpl) Delete(ctx context.Context, key string) error {
	err := global.Redis.Del(ctx, key).Err()
	if err != nil {
		r.log.LogError("Lỗi khi xóa key từ Redis.", err, zap.String("key", key))
		return err
	}

	r.log.LogInfo("Delete key", "Thực hiện thành công.", zap.String("key", key))
	return nil
}

func (r *RedisRepositoryImpl) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := global.Redis.Exists(ctx, key).Result()
	if err != nil {
		r.log.LogError("Lỗi khi kiểm tra tồn tại key trong Redis.", err, zap.String("key", key))
		return false, err
	}
	return exists > 0, nil
}

func (r *RedisRepositoryImpl) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := global.Redis.Expire( ctx, key, expiration).Err()
	if err != nil {
		r.log.LogError("Lỗi khi đặt thời gian hết hạn cho key trong Redis.", err, zap.String("key", key), zap.Duration("expiration", expiration))
		return err
	}
	r.log.LogInfo("Expire key", "Thực hiện thành công.", zap.String("key", key), zap.Duration("expiration", expiration))
	return nil
}