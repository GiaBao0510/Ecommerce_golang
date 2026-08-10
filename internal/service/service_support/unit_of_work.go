package servicesupport

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)

func WithTransaction(
	ctx context.Context, 	// Context dùng để quản lý vòng đời của transaction và truyền thông tin liên quan đến request
	db *sql.DB, 			// Cơ sở dữ liệu SQL để bắt đầu transaction
	q *database.Queries, 	// Instance của Queries để thực hiện các lệnh SQL trong transaction
	logger *zap.Logger, 	// Logger để ghi log thông tin liên quan đến transaction
	fn func(q *database.Queries) error, // Hàm callback chứa các lệnh SQL cần thực hiện trong transaction
) error {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Error("WithTransaction: Không thể bắt đầu transaction", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	// Tạo một instance mới của Queries với transaction
	// Mọi lệnh SQL qua qtx (thay vì q gốc) bên tron fn nằm chung 1 transaction
	qtx := q.WithTx(tx)

	if err := fn(qtx); err != nil {
		// Xuất hiện lỗi tại bất kỳ bước nào đó trong fn -> hủy toàn bộ thay đổi
		if rbErr := tx.Rollback(); rbErr != nil {
			logger.Error("WithTransaction: Không thể rollback transaction", zap.Error(rbErr), zap.NamedError("original_error", err))
		}

		// trả về lỗi gốc
		return err
	}

	// Nếu tất cả các lệnh SQL trong fn thành công -> commit transaction
	if err := tx.Commit(); err != nil {
		logger.Error("WithTransaction: Không thể commit transaction", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	return nil
}
