package repository

import (
	"context"
	"time"
)

type IRedisRepository interface {

	// Key - value (type string) để lưu trữ dữ liệu trong Redis. Key là chuỗi định danh duy nhất, còn value là dữ liệu được lưu trữ dưới dạng chuỗi.
	Set(ctx context.Context, key, value string, expiration time.Duration) error // Phương thức này cho phép lưu trữ một cặp key-value trong Redis với thời gian hết hạn (expiration) được chỉ định. Nếu expiration là 0, key sẽ tồn tại mãi mãi cho đến khi bị xóa.
	Get(ctx context.Context, key string) (string, error)                        // Phương thức này cho phép truy xuất giá trị (value) từ Redis dựa trên key. Nếu key không tồn tại, nó sẽ trả về lỗi.
	Delete(ctx context.Context, key string) error                               // Phương thức này cho phép xóa một key khỏi Redis. Nếu key không tồn tại, nó sẽ trả về lỗi.
	Exists(ctx context.Context, key string) (bool, error)                       // Phương thức này kiểm tra xem một key có tồn tại trong Redis hay không. Nó trả về true nếu key tồn tại, ngược lại trả về false. Nếu có lỗi trong quá trình kiểm tra, nó sẽ trả về lỗi đó.
	Expire(ctx context.Context, key string, expiration time.Duration) error     // Phương thức này cho phép đặt thời gian hết hạn (expiration) cho một key trong Redis. Nếu expiration là 0, key sẽ tồn tại mãi mãi. Nếu key không tồn tại, nó sẽ trả về lỗi.

	// Key - value (type )
}
