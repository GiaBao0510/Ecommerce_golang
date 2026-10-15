package repositoryimpl

import (
	"context"
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type BrandRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewBrandRepository(db *database.Queries, logger *zap.Logger) repository.IBrandRepository {
	return &BrandRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "BrandRepository"),
	}
}

func (r *BrandRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	rows, err := r.db.GetAllBrands(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}
	if len(rows) == 0 {
		r.dblog.LogWarning("GetAll", "No brands found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy brand nào")
	}
	result := make([]models.Brand, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapper.ToBrandModel(row))
	}
	return result, nil
}

func (r *BrandRepository) GetByID(ctx context.Context, id int32) (*models.Brand, error) {
	row, err := r.db.GetBrandByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy brand với ID: " + strconv.Itoa(int(id)))
	}
	result := mapper.ToBrandModel(row)
	return &result, nil
}

func (r *BrandRepository) GetByName(ctx context.Context, name string) (*models.Brand, error) {
	row, err := r.db.GetBrandByName(ctx, name)
	if err != nil {
		r.dblog.LogError("GetByName", err, zap.String("name", name))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy brand với tên: " + name)
	}
	result := mapper.ToBrandModel(row)
	return &result, nil
}

func (r *BrandRepository) Create(ctx context.Context, obj *models.Brand) error {
	if err := r.db.CreateBrand(ctx, obj.Name); err != nil {
		r.dblog.LogError("Create", err, zap.String("name", obj.Name))
		return apperrors.NewInternalServerError(err)
	}
	return nil
}

func (r *BrandRepository) Delete(ctx context.Context, id int32) error {
	result, err := r.db.DeleteBrand(ctx, id)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Delete", "Không tìm thấy brand với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *BrandRepository) Update_Put(ctx context.Context, id int32, obj *models.Brand) error {
	result, err := r.db.UpdateBrandByID_PUT(ctx, database.UpdateBrandByID_PUTParams{
		Name:    obj.Name,
		IDBrand: id,
	})
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Update_Put", "Không tìm thấy brand với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}
