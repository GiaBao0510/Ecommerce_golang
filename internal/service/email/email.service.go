package email

import (
	"context"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type IEmailService interface {
	// SendNotification gửi thông báo thường đến email người dùng
	SendNotification(ctx context.Context, email, message string) error
	// SendOTP tạo mã OTP, lưu vào Redis, gửi email chứa OTP cho người dùng
	SendOTP(ctx context.Context, email string) error
}

// Thông qua IEmailProviderService, EmailService có thể gửi email thông qua nhà cung cấp email cụ thể
type EmailService struct {
	provider IEmailProviderService
	userRepo repository.IUserRepository
	redisRepo repository.IRedisRepository
	logger *loghelper.ServiceLogger
}

// Khởi tạo EmailService với dependency injection của IEmailProviderService, IUserRepository, IRedisRepository và ServiceLogger
func NewEmailService(
	provider IEmailProviderService,
	userRepo repository.IUserRepository,
	redisRepo repository.IRedisRepository,
	logger *loghelper.ServiceLogger,
) IEmailService {
	return &EmailService{
		provider: provider,
		userRepo: userRepo,
		redisRepo: redisRepo,
		logger: logger,
	}
}


// Gửi email thông báo
func(s *EmailService) SendNotification(ctx context.Context, email, message string) error {
	
	// Kiểm tra người dùng có tồn tại trong cơ sở dữ liệu hay không
	_, err := s.isUserExists(ctx, email)
	if err != nil {
		return err
	}

	// Gửi thông báo cho người dùng
	data := dto.EmailMessage{
		To: email,
		Subject: "Thông báo từ hệ thống",
		Text: message,
		Category: "notification",
	}

	return s.provider.Send(ctx, &data)
}

//Gửi email với mã OTP để xác thực người dùng
// -------------------------------------------
// SendOTP tạo mã OTP 6 chữ số, lưu vào Redis (hết hạn sau 5 phút),
// rồi gửi email chứa mã OTP đến người dùng.
//
// Flow:
//   1. Kiểm tra email có tồn tại trong DB không
//   2. Kiểm tra user đã verify chưa (nếu rồi → báo lỗi)
//   3. Tạo OTP ngẫu nhiên 6 số
//   4. Lưu OTP vào Redis với key = "otp:{email}", TTL = 5 phút
//   5. Gửi email chứa OTP qua provider
func (s *EmailService) SendOTP(
    ctx context.Context,
    email string,
) error {
	
	// Kiểm tra người dùng có tồn tại trong cơ sở dữ liệu hay không
	status, err := s.isUserExists(ctx, email)
	if err != nil {
		return err
	}

	// Nếu người dùng đã tồn tại mà đã được xác thực, không cần gửi mã OTP
	if status == 1 {
		s.logger.LogInfo("User already verified, no need to send OTP: ", "",
			zap.String("to_email", email),
		)
		return apperrors.NewBadRequestError("User already verified, no need to send OTP")
	}

	// Tạo mã OTP ngẫu nhiên 6 chữ số
	otp := util.GenerateRandomNumber(6)

	// Lưu OTP vào Redis với key = "otp:{email}", TTL = 5 phút
	if err := s.redisRepo.Set(ctx, "otp:" + email, otp, 5 * time.Minute); err != nil {
		s.logger.LogError("Failed to store OTP in Redis: ", err)
		return apperrors.NewDetailedInternalServerError("Failed to store OTP in Redis", err)
	}

	data := dto.EmailMessage{
		To: email,
		Subject: "Mã xác thực từ hệ thống",
		Text: util.OTP_CodeSendingTemplate(otp),
		Category: "otp_verification",
	}

	return s.provider.Send(ctx, &data)
}


// Hàm kiểm tra người dùng có tồn tại trong cơ sở dữ liệu hay không
func(s *EmailService) isUserExists(ctx context.Context, email string) (int, error) {
	result, err := s.userRepo.UserEmailVerificationStatus(ctx, email)
	if err != nil {
		s.logger.LogError("Failed to check user email verification status: ", err)
		return -1,apperrors.NewDetailedInternalServerError("Failed to check user email verification status", err)
	}

	if result < 0 {
		s.logger.LogError("Email does not exist in the database: ", nil)
		return -1, apperrors.NewNotFoundError("Email does not exist in the database")
	}

	return result, nil
}