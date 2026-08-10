package authen

import (
	"context"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type VerifyUserUsecase struct {
	userRepo  repository.IUserRepository
	redisRepo repository.IRedisRepository
	emailRepo repository.IEmailRepository
	logger    *loghelper.DBLogger
}

// Khởi động
func NewVerifyUserUsecase(
	userRepo repository.IUserRepository,
	emailRepo repository.IEmailRepository,
	redisRepo repository.IRedisRepository,
	logger *loghelper.DBLogger,
) VerifyUserUsecase {
	return VerifyUserUsecase{
		userRepo:  userRepo,
		redisRepo: redisRepo,
		logger:    logger,
		emailRepo: emailRepo,
	}
}

// Thực hiện gửi email xác thực cho người dùng
func (u *VerifyUserUsecase) SendVerificationEmail(ctx context.Context, email string) error {

	// 1. Kiểm tra xem email người dùng đã tồn tại không
	exists, err := u.userRepo.UserEmailExists(ctx, email)
	if err != nil {
		return err
	}

	if !exists {
		u.logger.LogWarning("SendVerificationEmail", "Email người dùng không tồn tại trong cơ sở dữ liệu", zap.String("email", email))
		return apperrors.NewBadRequestError("Email người dùng không tồn tại")
	}

	// 2. kiểm tra xem email có tồn tại trong cơ sở dữ liệu không
	isNotVerified, err := u.userRepo.CheckUserEmailExists_HasNotBeenVerified(ctx, email)
	if err != nil {
		return err
	}

	// 3. Nếu tồn tại và đã xác thực rồi
	if !isNotVerified {
		u.logger.LogInfo("VerifyEmail", "Email người dùng đã được xác thực trước đó", zap.String("email", email))
		return nil
	}

	// 4. Nếu tồn tại mà chưa xác thực, thực hiện các bước xác thực email
	// 4.1 Tạo mã OTP
	otp := util.GenerateRandomNumber(6) // Tạo mã OTP gồm 6 chữ số

	// 4.2 Lưu mã OTP vào Redis với thời gian hết hạn (5 phút)
	if err := u.redisRepo.Set(ctx, "otp:"+email, otp, 5*time.Minute); err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi lưu OTP vào Redis", err, zap.Error(err), zap.String("email", email))
		return err
	}

	// 4.3 Tạo nội dung email với mã OTP
	body := util.OTP_CodeSendingTemplate(otp)

	// 4.4 Tạo data lưu thông tin
	emailData := models.EmailData{
		ToEmail:  email,
		ToName:   "Client",
		Subject:  "Xác thực email của bạn",
		HTMLBody: body,
		TextBody: "Xác thực email của bạn với mã OTP: " + otp,
	}

	// 4.5 Gửi email xác thực đến người dùng
	if err := u.emailRepo.SendEmail(ctx, emailData); err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi gửi email xác thực", err, zap.Error(err), zap.String("email", email))
		return err
	}

	u.logger.LogInfo("VerifyEmail", "Email xác thực đã được gửi thành công", zap.String("email", email))
	return nil
}

// Thực hiện xác thực email người dùng bằng mã OTP
func (u *VerifyUserUsecase) VerifyEmail(ctx context.Context, email, otp string) error {

	// 1. Lấy mã OTP từ Redis
	storedOtp, err := u.redisRepo.Get(ctx, "otp:"+email)
	if err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi lấy OTP từ Redis", err, zap.Error(err), zap.String("email", email))
		return err
	}

	// 2. Xác minh mã OTP
	if storedOtp != otp {
		u.logger.LogWarning("VerifyEmail", "Mã OTP không hợp lệ", zap.String("email", email), zap.String("provided_otp", otp))
		return apperrors.NewBadRequestError("Mã OTP không hợp lệ")
	}

	// 3. Cập nhật trong cơ sở dữ liệu để đánh dấu email là đã xác thực
	if err := u.userRepo.VerifyUserEmail(ctx, email); err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi cập nhật trạng thái xác thực email trong cơ sở dữ liệu", err, zap.Error(err), zap.String("email", email))
		return err
	}

	// 4. Xóa mã OTP khỏi Redis sau khi xác thực thành công. Tại đây, nếu xóa thất bại, chúng ta chỉ log lỗi mà không trả về lỗi, vì xác thực đã thành công.
	if err := u.redisRepo.Delete(ctx, "otp:"+email); err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi xóa OTP khỏi Redis", err, zap.Error(err), zap.String("email", email))
	}

	return nil
}

// Thực hiện xác thực mã OTP qua email cho người dùng
func (u *VerifyUserUsecase) VerifyOTP_viaEmail(ctx context.Context, email, otp string) error {
	// 1. Lấy mã OTP từ Redis
	storedOtp, err := u.redisRepo.Get(ctx, "otp:"+email)
	if err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi lấy OTP từ Redis", err, zap.Error(err), zap.String("email", email))
		return err
	}

	// 2. Xác minh mã OTP
	if storedOtp != otp {
		u.logger.LogWarning("VerifyEmail", "Mã OTP không hợp lệ", zap.String("email", email), zap.String("provided_otp", otp))
		return apperrors.NewBadRequestError("Mã OTP không hợp lệ")
	}

	// 3. Xóa mã OTP khỏi Redis sau khi xác thực thành công. Tại đây, nếu xóa thất bại, chúng ta chỉ log lỗi mà không trả về lỗi, vì xác thực đã thành công.
	if err := u.redisRepo.Delete(ctx, "otp:"+email); err != nil {
		u.logger.LogError("Error[VerifyEmail]: Lỗi khi xóa OTP khỏi Redis", err, zap.Error(err), zap.String("email", email))
	}

	return nil
}
