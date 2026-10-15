package repositoryimpl

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type ProductDetailRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewProductDetailRepository(db *database.Queries, logger *zap.Logger) repository.IProductDetailRepository {
	return &ProductDetailRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "ProductDetailRepository"),
	}
}

func (r *ProductDetailRepository) GetAll(ctx context.Context) ([]models.ProductDetail, error) {
	rows, err := r.db.GetAllProductDetails(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}
	if len(rows) == 0 {
		r.dblog.LogWarning("GetAll", "No product details found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy product detail nào")
	}
	result := make([]models.ProductDetail, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapper.ToProductDetailModel(row))
	}
	return result, nil
}

func (r *ProductDetailRepository) GetByID(ctx context.Context, idProduct string) (*models.ProductDetail, error) {
	row, err := r.db.GetProductDetailByID(ctx, idProduct)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.String("id_product", idProduct))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy product detail với ID: " + idProduct)
	}
	result := mapper.ToProductDetailModel(row)
	return &result, nil
}

func (r *ProductDetailRepository) Create(ctx context.Context, obj *models.ProductDetail) error {
	params := database.CreateProductDetailParams{
		IDProduct:      obj.IDProduct,
		Price:          obj.Price,
		LikeCount:      obj.LikeCount,
		Quantity:       obj.Quantity,
		ProductionDate: obj.ProductionDate,
		ExpirationDate: obj.ExpirationDate,
	}
	if err := r.db.CreateProductDetail(ctx, params); err != nil {
		r.dblog.LogError("Create", err, zap.String("id_product", obj.IDProduct))
		return MapDBErrorWithContext(
			err,
			"Product không tồn tại hoặc dữ liệu product detail không hợp lệ",
		)
	}
	return nil
}

func (r *ProductDetailRepository) Delete(ctx context.Context, idProduct string) error {
	result, err := r.db.DeleteProductDetail(ctx, idProduct)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.String("id_product", idProduct))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Delete", "Không tìm thấy product detail với ID: "+idProduct, r.dblog, zap.String("id_product", idProduct))
}

func (r *ProductDetailRepository) Update_Put(ctx context.Context, idProduct string, obj *models.ProductDetail) error {
	result, err := r.db.UpdateProductDetailByID_PUT(ctx, database.UpdateProductDetailByID_PUTParams{
		Price:          obj.Price,
		LikeCount:      obj.LikeCount,
		Quantity:       obj.Quantity,
		ProductionDate: obj.ProductionDate,
		ExpirationDate: obj.ExpirationDate,
		IDProduct:      idProduct,
	})
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.String("id_product", idProduct))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Update_Put", "Không tìm thấy product detail với ID: "+idProduct, r.dblog, zap.String("id_product", idProduct))
}

func (r *ProductDetailRepository) UpdatePrice_Patch(ctx context.Context, idProduct string, price string) error {
	result, err := r.db.UpdateProductPriceByID_PATCH(ctx, database.UpdateProductPriceByID_PATCHParams{Price: price, IDProduct: idProduct})
	if err != nil {
		r.dblog.LogError("UpdatePrice_Patch", err, zap.String("id_product", idProduct))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdatePrice_Patch", "Không tìm thấy product detail với ID: "+idProduct, r.dblog, zap.String("id_product", idProduct))
}

func (r *ProductDetailRepository) UpdateLikeCount_Patch(ctx context.Context, idProduct string, likeCount int32) error {
	result, err := r.db.UpdateProductLikeCountByID_PATCH(ctx, database.UpdateProductLikeCountByID_PATCHParams{LikeCount: likeCount, IDProduct: idProduct})
	if err != nil {
		r.dblog.LogError("UpdateLikeCount_Patch", err, zap.String("id_product", idProduct))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdateLikeCount_Patch", "Không tìm thấy product detail với ID: "+idProduct, r.dblog, zap.String("id_product", idProduct))
}

func (r *ProductDetailRepository) UpdateQuantity_Patch(ctx context.Context, idProduct string, quantity int32) error {
	result, err := r.db.UpdateProductQuantityByID_PATCH(ctx, database.UpdateProductQuantityByID_PATCHParams{Quantity: quantity, IDProduct: idProduct})
	if err != nil {
		r.dblog.LogError("UpdateQuantity_Patch", err, zap.String("id_product", idProduct))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdateQuantity_Patch", "Không tìm thấy product detail với ID: "+idProduct, r.dblog, zap.String("id_product", idProduct))
}
