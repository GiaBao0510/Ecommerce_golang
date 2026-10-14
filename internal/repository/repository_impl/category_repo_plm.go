package repositoryimpl

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type CategoryRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewCategoryRepository(db *database.Queries, logger *zap.Logger) repository.ICategoryRepository {
	return &CategoryRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "CategoryRepository"),
	}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	rows, err := r.db.GetAllCategories(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}
	if len(rows) == 0 {
		r.dblog.LogWarning("GetAll", "No categories found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy category nào")
	}

	result := make([]models.Category, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapper.ToCategoryModel(row))
	}
	return result, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int32) (*models.Category, error) {
	row, err := r.db.GetCategoryByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy category với ID: " + strconv.Itoa(int(id)))
	}
	result := mapper.ToCategoryModel(row)
	return &result, nil
}

func (r *CategoryRepository) GetByName(ctx context.Context, name string) (*models.Category, error) {
	row, err := r.db.GetCategoryByName(ctx, name)
	if err != nil {
		r.dblog.LogError("GetByName", err, zap.String("name", name))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy category với tên: " + name)
	}
	result := mapper.ToCategoryModel(row)
	return &result, nil
}

func (r *CategoryRepository) Create(ctx context.Context, obj *models.Category) (int, error) {
	params := database.CreateCategoryParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
	}
	if err := r.db.CreateCategory(ctx, params); err != nil {
		r.dblog.LogError("Create", err, zap.String("name", obj.Name))
		return 0, apperrors.NewInternalServerError(err)
	}
	return 0, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int32) error {
	result, err := r.db.DeleteCategory(ctx, id)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Delete", "Không tìm thấy category với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *CategoryRepository) Update_Put(ctx context.Context, id int32, obj *models.Category) error {
	params := database.UpdateCategoryByID_PUTParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		IDCate: id,
	}
	result, err := r.db.UpdateCategoryByID_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Update_Put", "Không tìm thấy category với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}
