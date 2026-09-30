package authen

import (
	"net/http"
	"time"

	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service/oauth2"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/oauth2"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Flow oauth2 luôn có 2 nửa
//1. Nửa đầu sẽ redirect đến trang xác thực của provider (ở đây là Google) để người dùng đăng nhập và cấp quyền truy cập.
//2. Nửa sau sẽ nhận callback từ provider với mã xác thực, sau đó trao đổi mã này lấy access token và thông tin người dùng.

type LoginGoogleController struct {
	redisRepo    repository.IRedisRepository
	providerCtx  *oauth2.OAuth2ServiceContext
	loginUseCase *service.LoginWithGoogleUseCase
}

func NewLoginGoogleController(
	redisRepo repository.IRedisRepository,
	providerCtx *oauth2.OAuth2ServiceContext,
	loginUseCase *service.LoginWithGoogleUseCase,
) *LoginGoogleController {
	return &LoginGoogleController{
		redisRepo:    redisRepo,
		providerCtx:  providerCtx,
		loginUseCase: loginUseCase,
	}
}

// Login_Google: GET /authen/google/login
// Chuyển hướng người dùng sang trang chọn tài khoản Google
func (L *LoginGoogleController) Login_Google(ctx *gin.Context) error {
	state := uuid.New().String() // Tạo một state ngẫu nhiên để bảo vệ chống lại các cuộc tấn công CSRF

	if err := L.redisRepo.Set(ctx, "oauth2:"+state, "1", 5*time.Minute); err != nil {
		return controller.HandleValidationError(err)
	}

	ctx.Redirect(http.StatusTemporaryRedirect, L.providerCtx.GetAuthURL(state))
	return nil
}

// Login_GoogleCallback: GET /authen/google/callback
// Xử lý callback từ Google sau khi người dùng đăng nhập và cấp quyền truy cập
func (L *LoginGoogleController) Login_GoogleCallback(ctx *gin.Context) error {
	code := ctx.Query("code")
	state := ctx.Query("state")

	// Kiểm tra xem state có tồn tại trong Redis không
	if code == "" || state == "" {
		return controller.HandleValidationError(apperrors.NewBadRequestError("Missing code or state in callback"))
	}

	// Kiểm tra state trong Redis để bảo vệ chống lại các cuộc tấn công CSRF
	exists, err := L.redisRepo.Get(ctx, "oauth2:"+state)
	if err != nil || exists == "" {
		return controller.HandleValidationError(apperrors.NewBadRequestError("Invalid or expired state"))
	}

	// Xóa state khỏi Redis sau khi đã sử dụng
	_ = L.redisRepo.Delete(ctx, "oauth2:"+state)

	// Trao đổi mã code lấy access token và thông tin người dùng
	token, err := L.providerCtx.ExchangeCode(ctx, code)
	if err != nil {
		return controller.HandleValidationError(err)
	}

	info, err := L.providerCtx.GetUserInfo(ctx, token)
	if err != nil {
		return controller.HandleValidationError(err)
	}
	if info.Email == "" {
		return controller.HandleValidationError(apperrors.NewBadRequestError("Email not found in user info"))
	}

	result, err := L.loginUseCase.Login(ctx, &models.CreateUsersRequestNonStrict{
		Email:      info.Email,
		User_name:  info.Name,
		Avatar_url: info.AvatarURL,
	})

	if err != nil {
		return controller.HandleValidationError(err)
	}

	response.Success_Response(ctx, http.StatusOK, "Đăng nhập Google thành công", result)
	return nil
}
