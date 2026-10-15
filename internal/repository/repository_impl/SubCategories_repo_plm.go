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

type SubCategoryRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewSubCategoryRepository(db *database.Queries, logger *zap.Logger) repository.ISubCategoryRepository {
	return &SubCategoryRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "SubCategoryRepository"),
	}
}

func (r *SubCategoryRepository) GetAll(ctx context.Context) ([]models.SubCategory, error) {
	rows, err := r.db.GetAllSubCategories(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}
	if len(rows) == 0 {
		r.dblog.LogWarning("GetAll", "No subcategories found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy subcategory nào")
	}

	result := make([]models.SubCategory, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapper.ToSubCategoryModel(row))
	}
	return result, nil
}

func (r *SubCategoryRepository) GetByID(ctx context.Context, id int32) (*models.SubCategory, error) {
	row, err := r.db.GetSubCategoryByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy subcategory với ID: " + strconv.Itoa(int(id)))
	}
	result := mapper.ToSubCategoryModel(row)
	return &result, nil
}

func (r *SubCategoryRepository) GetByName(ctx context.Context, name string) (*models.SubCategory, error) {
	row, err := r.db.GetSubCategoryByName(ctx, name)
	if err != nil {
		r.dblog.LogError("GetByName", err, zap.String("name", name))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy subcategory với tên: " + name)
	}
	result := mapper.ToSubCategoryModel(row)
	return &result, nil
}

func (r *SubCategoryRepository) Create(ctx context.Context, obj *models.SubCategory) (int, error) {
	params := database.CreateSubCategoryParams{
		IDCate: obj.IDCate,
		Name:   obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
	}
	if err := r.db.CreateSubCategory(ctx, params); err != nil {
		r.dblog.LogError("Create", err, zap.String("name", obj.Name), zap.Int32("category_id", obj.IDCate))
		return 0, MapDBErrorWithContext(
			err,
			"Category không tồn tại hoặc dữ liệu subcategory không hợp lệ",
		)
	}
	return 0, nil
}

func (r *SubCategoryRepository) Delete(ctx context.Context, id int32) error {
	result, err := r.db.DeleteSubCategory(ctx, id)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Delete", "Không tìm thấy subcategory với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *SubCategoryRepository) Update_Put(ctx context.Context, id int32, obj *models.SubCategory) error {
	params := database.UpdateSubCategoryByID_PUTParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		IDCate: id,
	}
	result, err := r.db.UpdateSubCategoryByID_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Update_Put", "Không tìm thấy subcategory với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}
