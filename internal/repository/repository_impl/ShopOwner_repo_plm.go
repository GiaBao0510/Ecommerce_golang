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

type ShopOwnerRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewShopOwnerRepository(db *database.Queries, logger *zap.Logger) repository.IShopOwnerRepository {
	return &ShopOwnerRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "ShopOwnerRepository"),
	}
}

func (r *ShopOwnerRepository) GetAll(ctx context.Context) ([]models.ShopOwner, error) {
	rows, err := r.db.GetAllShopOwners(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}
	if len(rows) == 0 {
		r.dblog.LogWarning("GetAll", "No shop owners found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy shop owner nào")
	}
	result := make([]models.ShopOwner, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapper.ToShopOwnerModel(row))
	}
	return result, nil
}

func (r *ShopOwnerRepository) GetByID(ctx context.Context, id int32) (*models.ShopOwner, error) {
	row, err := r.db.GetShopOwnerByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy shop owner với ID: " + strconv.Itoa(int(id)))
	}
	result := mapper.ToShopOwnerModel(row)
	return &result, nil
}

func (r *ShopOwnerRepository) GetByName(ctx context.Context, name string) (*models.ShopOwner, error) {
	row, err := r.db.GetShopOwnerByName(ctx, name)
	if err != nil {
		r.dblog.LogError("GetByName", err, zap.String("name", name))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy shop owner với tên: " + name)
	}
	result := mapper.ToShopOwnerModel(row)
	return &result, nil
}

func (r *ShopOwnerRepository) Create(ctx context.Context, obj *models.ShopOwner) error {
	params := database.CreateShopOwnerParams{
		Uuid: obj.Uuid,
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		AvatarUrl: sql.NullString{
			String: obj.AvatarUrl,
			Valid:  obj.AvatarUrl != "",
		},
	}
	if err := r.db.CreateShopOwner(ctx, params); err != nil {
		r.dblog.LogError("Create", err, zap.String("uuid", obj.Uuid), zap.String("name", obj.Name))
		return apperrors.NewInternalServerError(err)
	}
	return nil
}

func (r *ShopOwnerRepository) Delete(ctx context.Context, id int32) error {
	result, err := r.db.DeleteShopOwner(ctx, id)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Delete", "Không tìm thấy shop owner với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *ShopOwnerRepository) Update_Put(ctx context.Context, id int32, obj *models.ShopOwner) error {
	result, err := r.db.UpdateShopOwnerByID_PUT(ctx, database.UpdateShopOwnerByID_PUTParams{
		Uuid: obj.Uuid,
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		AvatarUrl: sql.NullString{
			String: obj.AvatarUrl,
			Valid:  obj.AvatarUrl != "",
		},
		IDShopOwner: id,
	})
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "Update_Put", "Không tìm thấy shop owner với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *ShopOwnerRepository) UpdateName_Patch(ctx context.Context, id int32, name string) error {
	result, err := r.db.UpdateShopOwnerNameByID_PATCH(ctx, database.UpdateShopOwnerNameByID_PATCHParams{Name: name, IDShopOwner: id})
	if err != nil {
		r.dblog.LogError("UpdateName_Patch", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdateName_Patch", "Không tìm thấy shop owner với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *ShopOwnerRepository) UpdateDescription_Patch(ctx context.Context, id int32, description string) error {
	result, err := r.db.UpdateShopOwnerDescriptionByID_PATCH(ctx, database.UpdateShopOwnerDescriptionByID_PATCHParams{
		Description: sql.NullString{String: description, Valid: description != ""},
		IDShopOwner: id,
	})
	if err != nil {
		r.dblog.LogError("UpdateDescription_Patch", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdateDescription_Patch", "Không tìm thấy shop owner với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}

func (r *ShopOwnerRepository) UpdateAvatar_Patch(ctx context.Context, id int32, avatarURL string) error {
	result, err := r.db.UpdateShopOwnerAvatarByID_PATCH(ctx, database.UpdateShopOwnerAvatarByID_PATCHParams{
		AvatarUrl:   sql.NullString{String: avatarURL, Valid: avatarURL != ""},
		IDShopOwner: id,
	})
	if err != nil {
		r.dblog.LogError("UpdateAvatar_Patch", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}
	return CheckRowsAffected(result, "UpdateAvatar_Patch", "Không tìm thấy shop owner với ID: "+strconv.Itoa(int(id)), r.dblog, zap.Int32("id", id))
}
