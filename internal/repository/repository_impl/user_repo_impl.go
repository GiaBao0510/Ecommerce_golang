package repositoryimpl

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	dto "github.com/GiaBao0510/Ecommerce_golang/internal/dto/user"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

// triển khai
func NewUserRepository(db *database.Queries, logger *zap.Logger) repository.IUserRepository {
	return &UserRepository{db: db, dblog: loghelper.NewDBLogger(logger, "UserRepository")}
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.Users, error) {
	rows, err := r.db.GetUserByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.String("id", id))
		return nil, MapDBErrorWithContext(err, "Không tìm thấy người dùng với ID: "+id)
	}

	result := mapper.ToUserModel(rows)
	return &result, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	rows, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		r.dblog.LogError("GetUserByEmail", err, zap.String("email", email))
		return nil, MapDBErrorWithContext(err, "Không tìm thấy người dùng với email: "+email)
	}

	result := mapper.ToUserModel(rows)
	return &result, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error) {

	rows, err := r.db.GetUserByPhone(ctx, phone)
	if err != nil {
		r.dblog.LogError("GetUserByPhone", err, zap.String("phone", phone.String))
		return nil, MapDBErrorWithContext(err, "Không tìm thấy người dùng với điện thoại: "+phone.String)
	}

	result := mapper.ToUserModel(rows)
	return &result, nil
}

func (r *UserRepository) GetUID_PasswordHashByEmail(ctx context.Context, email string) (*dto.UserResponseBase, error) {
	row, err := r.db.GetUID_PasswordHashByEmail(ctx, email)
	if err != nil {
		r.dblog.LogError("GetUID_PasswordHashByEmail", err, zap.String("email", email))
		return nil, MapDBErrorWithContext(err, "Không tìm thấy người dùng với email: "+email)
	}

	result := dto.UserResponseBase{
		Uuid:          row.Uuid,
		Password_hash: row.PasswordHash,
	}
	return &result, nil
}

func (r *UserRepository) GetUID_PasswordHashByPhone(ctx context.Context, phone sql.NullString) (*dto.UserResponseBase, error) {
	row, err := r.db.GetUID_PasswordHashByPhone(ctx, phone)
	if err != nil {
		r.dblog.LogError("GetUID_PasswordHashByPhone", err, zap.String("phone", phone.String))
		return nil, MapDBErrorWithContext(err, "Không tìm thấy người dùng với điện thoại: "+phone.String)
	}

	result := dto.UserResponseBase{
		Uuid:          row.Uuid,
		Password_hash: row.PasswordHash,
	}
	return &result, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.Users, error) {
	query, err := r.db.GetAllUsers(ctx)
	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, MapDBErrorWithContext(err, "Lỗi khi lấy danh sách người dùng")
	}

	if len(query) == 0 {
		r.dblog.LogWarning("GetAll", "No users found")
		return nil, MapDBErrorWithContext(apperrors.NewNotFoundError("Không tìm thấy người dùng nào"), "Không tìm thấy người dùng nào")
	}

	var users []models.Users
	for _, user := range query {
		users = append(users, mapper.ToUserModel(user))
	}

	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, obj *models.CreateUsersRequest) (string, error) {

	// Nhận các giá trị từ obj và chuẩn bị các tham số cho truy vấn SQL
	params := database.CreateUserParams{
		Uuid: uuid.NewString(), // Tạo UUID mới cho người dùng
		IDStatus: sql.NullInt32{
			Int32: obj.Id_status,
			Valid: obj.Id_status != 0,
		},
		UserName: obj.User_name,
		BirthDate: sql.NullTime{
			Time:  obj.Birth_date.Time,
			Valid: !obj.Birth_date.IsZero(),
		},
		Email: obj.Email,
		PhoneNum: sql.NullString{
			String: obj.Phone_num,
			Valid:  obj.Phone_num != "",
		},
		Address: sql.NullString{
			String: obj.Address,
			Valid:  obj.Address != "",
		},
		PasswordHash: obj.Password_hash,
		AvatarUrl: sql.NullString{
			String: obj.Avatar_url,
			Valid:  obj.Avatar_url != "",
		},
	}

	// Gọi phương thức CreateUser từ database.Queries để thực hiện việc tạo mới
	if err := r.db.CreateUser(ctx, params); err != nil {
		r.dblog.LogError("Create", err, zap.String("name", obj.User_name))
		return "", MapDBErrorWithContext(err, "Lỗi khi tạo người dùng mới")
	}

	return params.Uuid, nil
}

func (r *UserRepository) Update_Put(ctx context.Context, id string, obj *models.UpdateUsersPutRequest) error {
	params := database.UpdateUser_PUTParams{
		IDStatus: sql.NullInt32{
			Int32: obj.Id_status,
			Valid: obj.Id_status != 0,
		},
		UserName: obj.User_name,
		BirthDate: sql.NullTime{
			Time:  obj.Birth_date.Time,
			Valid: !obj.Birth_date.IsZero(),
		},
		Email: obj.Email,
		PhoneNum: sql.NullString{
			String: obj.Phone_num,
			Valid:  obj.Phone_num != "",
		},
		Address: sql.NullString{
			String: obj.Address,
			Valid:  obj.Address != "",
		},
		Uuid: id,
	}

	result, err := r.db.UpdateUser_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("UpdateUser_PUT", err, zap.String("id", id))
		return MapDBErrorWithContext(err, "Lỗi khi cập nhật người dùng với ID: "+id)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	affected, err := result.RowsAffected()
	if err != nil {
		r.dblog.LogError("RowsAffected", err, zap.String("id", id))
		return MapDBErrorWithContext(err, "Lỗi khi kiểm tra số lượng bản ghi bị ảnh hưởng cho người dùng với ID: "+id)
	}

	if affected == 0 {
		r.dblog.LogWarning("UpdateUser_PUT", "No rows affected", zap.String("id", id))
		return MapDBErrorWithContext(apperrors.NewNotFoundError("Không tìm thấy người dùng với ID: "+id), "Không tìm thấy người dùng với ID: "+id)
	}

	return nil
}

func (r *UserRepository) Update_Patch(ctx context.Context, id string, obj *models.UpdateUsersPatchRequest) error {
	var idStatus sql.NullInt32
	if obj.Id_status != nil {
		idStatus = sql.NullInt32{
			Int32: *obj.Id_status,
			Valid: *obj.Id_status != 0,
		}
	}

	var userName string
	if obj.User_name != nil {
		userName = *obj.User_name
	}

	var birthDate sql.NullTime
	if obj.Birth_date != nil {
		birthDate = sql.NullTime{
			Time:  obj.Birth_date.Time,
			Valid: !obj.Birth_date.IsZero(),
		}
	}

	var email string
	if obj.Email != nil {
		email = *obj.Email
	}

	var phoneNum sql.NullString
	if obj.Phone_num != nil {
		phoneNum = sql.NullString{
			String: *obj.Phone_num,
			Valid:  *obj.Phone_num != "",
		}
	}

	var address sql.NullString
	if obj.Address != nil {
		address = sql.NullString{
			String: *obj.Address,
			Valid:  *obj.Address != "",
		}
	}

	params := database.UpdateUser_PATCHParams{
		IDStatus:  idStatus,
		UserName:  userName,
		BirthDate: birthDate,
		Email:     email,
		PhoneNum:  phoneNum,
		Address:   address,
		Uuid:      id,
	}

	result, err := r.db.UpdateUser_PATCH(ctx, params)
	if err != nil {
		r.dblog.LogError("UpdateUser_PATCH", err, zap.String("id", id))
		return MapDBErrorWithContext(err, "Lỗi khi cập nhật người dùng với ID: "+id)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	affected, err := result.RowsAffected()
	if err != nil {
		r.dblog.LogError("RowsAffected", err, zap.String("id", id))
		return MapDBErrorWithContext(err, "Lỗi khi kiểm tra số lượng bản ghi bị ảnh hưởng cho người dùng với ID: "+id)
	}

	if affected == 0 {
		r.dblog.LogWarning("UpdateUser_PATCH", "No rows affected", zap.String("id", id))
		return MapDBErrorWithContext(apperrors.NewNotFoundError("Không tìm thấy người dùng với ID: "+id), "Không tìm thấy người dùng với ID: "+id)
	}

	return nil
}

func (r *UserRepository) UpdateUserPassword_PATCH(ctx context.Context, id string, passwordHash string) error {
	params := database.UpdateUserPassword_PATCHParams{
		PasswordHash: passwordHash,
		Uuid:         id,
	}

	result, err := r.db.UpdateUserPassword_PATCH(ctx, params)
	if err != nil {
		r.dblog.LogError("UpdateUserPassword_PATCH", err, zap.String("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	affected, err := result.RowsAffected()
	if err != nil {
		r.dblog.LogError("RowsAffected", err, zap.String("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		r.dblog.LogWarning("UpdateUser_PATCH", "No rows affected", zap.String("id", id))
		return apperrors.NewNotFoundError("Không tìm thấy người dùng với ID: " + id)
	}

	return nil
}
func (r *UserRepository) UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error {
	params := database.UpdateUserAvatar_PATCHParams{
		AvatarUrl: sql.NullString{
			String: avatarURL,
			Valid:  avatarURL != "",
		},
		Uuid: id,
	}

	result, err := r.db.UpdateUserAvatar_PATCH(ctx, params)
	if err != nil {
		r.dblog.LogError("UpdateUserAvatar_PATCH", err, zap.String("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if err := CheckRowsAffected(
		result,
		"UpdateUserAvatar_PATCH",
		"Không tìm thấy người dùng với ID: "+id,
		r.dblog,
		zap.String("id", id),
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.DeleteUser(ctx, id)
	if err != nil {
		r.dblog.LogError("DeleteUser", err, zap.String("id", id))
		return MapDBErrorWithContext(err, "Lỗi khi xóa người dùng với ID: "+id)
	}

	if err := CheckRowsAffected(
		result,
		"DeleteUser",
		"Không tìm thấy người dùng với ID: "+id,
		r.dblog,
		zap.String("id", id),
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) VerifyUserEmail(ctx context.Context, email string) error {
	result, err := r.db.VerifyEmail(ctx, email)
	if err != nil {
		r.dblog.LogError("VerifyUserEmail", err, zap.String("email", email))
		return MapDBErrorWithContext(err, "Lỗi khi xác thực email người dùng với email: "+email)
	}

	if err := CheckRowsAffected(
		result,
		"VerifyUserEmail",
		"Không tìm thấy người dùng với email: "+email,
		r.dblog,
		zap.String("email", email),
	); err != nil {
		return err
	}

	r.dblog.LogInfo("VerifyUserEmail", "Email người dùng đã được xác thực thành công", zap.String("email", email))

	return nil
}

func (r *UserRepository) VerifyUserPhone(ctx context.Context, phone string) error {
	result, err := r.db.VerifyPhone(ctx, sql.NullString{String: phone, Valid: true})
	if err != nil {
		r.dblog.LogError("VerifyUserPhone", err, zap.String("phone", phone))
		return MapDBErrorWithContext(err, "Lỗi khi xác thực số điện thoại người dùng với số điện thoại: "+phone)
	}

	if err := CheckRowsAffected(
		result,
		"VerifyUserPhone",
		"Không tìm thấy người dùng với số điện thoại: "+phone,
		r.dblog,
		zap.String("phone", phone),
	); err != nil {
		return err
	}

	r.dblog.LogInfo("VerifyUserPhone", "Số điện thoại người dùng đã được xác thực thành công", zap.String("phone", phone))

	return nil
}

func (r *UserRepository) CheckUserEmailExists_HasNotBeenVerified(ctx context.Context, email string) (bool, error) {
	result, err := r.db.UserEmailExists_HasNotBeenVerified(ctx, email)

	if err != nil {
		r.dblog.LogError("CheckUserEmailExists_HasNotBeenVerified", err, zap.String("email", email))
		return false, MapDBErrorWithContext(err, "Lỗi khi kiểm tra trạng thái xác thực email người dùng với email: "+email)
	}

	return result, nil
}

func (r *UserRepository) CheckUserPhoneExists_HasNotBeenVerified(ctx context.Context, phone string) (bool, error) {
	result, err := r.db.UserPhoneExists_HasNotBeenVerified(ctx, sql.NullString{String: phone, Valid: true})

	if err != nil {
		r.dblog.LogError("CheckUserPhoneExists_HasNotBeenVerified", err, zap.String("phone", phone))
		return false, MapDBErrorWithContext(err, "Lỗi khi kiểm tra trạng thái xác thực số điện thoại người dùng với số điện thoại: "+phone)
	}

	return result, nil
}

func (r *UserRepository) UserEmailExists(ctx context.Context, email string) (bool, error) {
	result, err := r.db.UserEmailExists(ctx, email)
	if err != nil {
		r.dblog.LogError("UserEmailExists", err, zap.String("email", email))
		return false, MapDBErrorWithContext(err, "Lỗi khi kiểm tra sự tồn tại của email người dùng với email: "+email)
	}

	return result, nil
}

func (r *UserRepository) UserPhoneExists(ctx context.Context, phone string) (bool, error) {
	phoneNull := sql.NullString{String: phone, Valid: true}
	result, err := r.db.UserPhoneExists(ctx, phoneNull)
	if err != nil {
		r.dblog.LogError("UserPhoneExists", err, zap.String("phone", phone))
		return false, MapDBErrorWithContext(err, "Lỗi khi kiểm tra sự tồn tại của số điện thoại người dùng với số điện thoại: "+phone)
	}

	return result, nil
}
