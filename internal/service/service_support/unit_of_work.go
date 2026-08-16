package servicesupport

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)

// hàm RunInTransaction. Hàm này là nền tảng để lo phần "Khung xương" của 1 transaction
// Tại đây Callback "fn" sẽ nhận về sql.Tx THUẦN. Và dùng khi 1 service cần phối hợp nhiều repository khác nhau trong cùng 1 transaction.
func RunInTx(
	ctx context.Context,
	db *sql.DB,
	logger *zap.Logger,
	fn func(tx *sql.Tx) error,
) error {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		logger.Error("RunInTx: Không thể bắt đầu transaction", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	// Thực hiện callback fn với transaction tx
	if err := fn(tx); err != nil{
		if rbErr := tx.Rollback(); rbErr != nil {
			logger.Error("RunInTx: Không thể rollback transaction", 
				zap.Error(rbErr), 
				zap.NamedError("original_error", err),
			)
		}
		return err
	}

	// Nếu tất cả các lệnh SQL trong fn thành công -> commit transaction
	if err := tx.Commit(); err != nil {
		logger.Error("RunInTx: Không thể commit transaction", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// Hàm WithTransaction thực hiện các thao tác trong một transaction. Nó nhận vào một hàm callback fn, trong đó chứa các lệnh SQL cần thực hiện.
func WithTransaction(
	ctx context.Context, // Context dùng để quản lý vòng đời của transaction và truyền thông tin liên quan đến request
	db *sql.DB, // Cơ sở dữ liệu SQL để bắt đầu transaction
	q *database.Queries, // Instance của Queries để thực hiện các lệnh SQL trong transaction
	logger *zap.Logger, // Logger để ghi log thông tin liên quan đến transaction
	fn func(q *database.Queries) error, // Hàm callback chứa các lệnh SQL cần thực hiện trong transaction
) error {

	// Phiên dịch từ sql.Tx sang database.Queries để có thể sử dụng các phương thức của Queries trong transaction
	return  RunInTx(ctx, db, logger, func(tx *sql.Tx) error {
		qtx := q.WithTx(tx)
		return  fn(qtx) 
	})
}