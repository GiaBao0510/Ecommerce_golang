package repositoryimpl

import (
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

// Hàm kiếm tra số hàng tác động
func CheckRowsAffected(result sql.Result, operation string, notFoundMessage string, logger *loghelper.DBLogger, logFields ...zap.Field) error {
	affected, err := result.RowsAffected()
	if err != nil {
		logger.LogError(operation, err , logFields...)
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		logger.LogWarning(operation, "No rows affected", logFields...)
		return apperrors.NewNotFoundError(notFoundMessage)
	}
	
	return nil
}