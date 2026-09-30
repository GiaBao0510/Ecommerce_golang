package oauth2

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// ══════════════════════════════════════════════════════════════════════════════
// Tầng 1: GoogleProviderStrategy, đây là chiến lược OAuth2 cụ thể cho Google, triển khai
// các phương thức cần thiết để lấy URL xác thực, trao đổi mã xác thực và lấy thông tin người dùng từ Google OAuth2 provider.
// ════════════════════════════════════════════════════════════════════════════════════
type GoogleProviderStrategy struct {
	config *oauth2.Config
}

func NewGoogleProviderStrategy() IOAuth2ServiceStrategy {
	return &GoogleProviderStrategy{
		config: &oauth2.Config{
			ClientID: global.Config.Authentication.OAuth2.Google.ClientID,
			ClientSecret: global.Config.Authentication.OAuth2.Google.ClientSecret,
			RedirectURL: global.Config.Authentication.OAuth2.Google.RedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},	
	}
}

// GetAuthURL trả về URL xác thực từ Google OAuth2 provider, sử dụng state để bảo vệ chống lại các cuộc tấn công CSRF.
func( g *GoogleProviderStrategy) GetAuthURL(state string) string {
	return g.config.AuthCodeURL(state ,oauth2.AccessTypeOffline)
}

// Hàm Exchange để trao đổi mã xác thực lấy được từ Google OAuth2 provider để lấy access token. Nó nhận vào mã xác thực và trả về access token hoặc lỗi nếu có.
func(g *GoogleProviderStrategy) Exchange(ctx context.Context, code string) (*oauth2.Token, error){
	return g.config.Exchange(ctx, code)
}

// Hàm GetUserInfor để lấy thông tin người dùng từ Google OAuth2 provider bằng access token. Nó nhận vào access token và trả về thông tin người dùng hoặc lỗi nếu có.
func (g *GoogleProviderStrategy) GetUserInfor(ctx context.Context, token *oauth2.Token) (*OAuth2UserInfo, error) {
	client := g.config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, apperrors.NewDetailedInternalServerError("Lỗi khi lấy thông tin người dùng từ Google", err)
	}
	defer resp.Body.Close() // Đóng body sau khi đọc xong để tránh rò rỉ tài nguyên

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, apperrors.NewDetailedInternalServerError("Lỗi khi lấy thông tin người dùng từ Google: "+string(body), errors.New(resp.Status))
	} 

	var raw struct {
		Email string	 `json:"email"`
		VerifiedEmail bool `json:"verified_email"`
		Name string		 `json:"name"`
		Picture string	 `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err  != nil {
		return nil, apperrors.NewDetailedInternalServerError("Lỗi khi giải mã thông tin người dùng từ Google", err)
	}

	return &OAuth2UserInfo{
		Email: raw.Email,
		Name: raw.Name,
		AvatarURL: raw.Picture,
		EmailVerified: raw.VerifiedEmail,
	}, nil
}

// ══════════════════════════════════════════════════════════════════════════════
// Tầng 2: Sẽ triển khai UseCase cho việc đăng nhập thông qua Google OAuth2, sử dụng GoogleProviderStrategy để thực hiện các bước xác thực và lấy thông tin người dùng.
// Nó sẽ kiểm tra xem người dùng đã tồn tại trong cơ sở dữ liệu chưa, nếu chưa thì tạo mới người dùng và sau đó trả về thông tin đăng nhập.
// ════════════════════════════════════════════════════════════════════════════════════
const defaultUserRoleID int32 = 2

type LoginWithGoogleUseCase struct {
	userRepo     repository.IUserRepository
	userRoleRepo repository.IUserRoleRepository
	redisRepo    repository.IRedisRepository
	slog         *loghelper.ServiceLogger
	db *sql.DB
	zapLogger *zap.Logger
}

func NewLoginWithGoogleUseCase(
	userRepo repository.IUserRepository,
	userRoleRepo repository.IUserRoleRepository,
	redisRepo repository.IRedisRepository,
	slog *loghelper.ServiceLogger,
	db *sql.DB,
	zapLogger *zap.Logger,
) *LoginWithGoogleUseCase {
	return &LoginWithGoogleUseCase{
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		redisRepo:    redisRepo,
		slog:         loghelper.NewServiceLogger(zapLogger, "LoginWithGoogleUseCase"),
		db: db,
		zapLogger: zapLogger,
	}
}

// triển khai đăng nhập thông qua Google OAuth2
func (l *LoginWithGoogleUseCase) Login(
	ctx context.Context,
	req *models.CreateUsersRequestNonStrict,
) (*models.LoginResponse, error) {

	// Kiểm tra xem email có trong DB không
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaEmail(ctx, req.Email)

	var uid string
	switch {

	// trong tường hợp đã tồn tại
	case err == nil:
		// Đã tồn tại mà chưa xác nhận (tức là đã đăng ký nhưng chưa xác minh email), thì trả về lỗi
		if !userVeriInfor.Is_email_verified {
			l.slog.LogWarning("Đăng nhập thông quan Google","Error: Vì tài khoản này đã được tạo trước đó nhưng chưa xác minh email, nên không thể đăng nhập thông qua Google OAuth2", zap.String("email", req.Email))
			return nil, apperrors.NewDetailedEmailDuplicateError("Tài khoản này đã được tạo trước đó nhưng chưa xác minh email, nên không thể đăng nhập thông qua Google OAuth2")
		} 
		uid = userVeriInfor.Uuid

	// Nếu không tìm thấy người dùng trong DB, thì tạo mới người dùng
	case errors.Is(err, apperrors.ErrNotFound):
		// tạo user với transaction, nếu có lỗi thì rollback
		uid, err = l.createUserWithTransaction(ctx, req)
		if err != nil {
			l.slog.LogError("Lỗi tạo mới người dùng từ OAuth2", err, zap.Error(err))
			return nil, apperrors.NewDetailedInternalServerError("Lỗi tạo mới người dùng từ OAuth2", err)
		}

	// Lỗi thật
	default:
		l.slog.LogError("Lỗi truy xuất người dùng từ DB", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi truy xuất người dùng từ DB", err)
	}

	// truy xuất Role của người dùng dựa trên Uid
	roleID, err := l.userRoleRepo.GetRoleIDByUserID(ctx, uid)
	if err != nil {
		l.slog.LogError("Lỗi truy xuất Role của người dùng", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi truy xuất Role của người dùng", err)
	}

	// Nếu có thì đăng nhập và trả về thông tin người dùng
	return GenerateAccessTokenAndRefreshToken(
		l.slog,
		l.redisRepo,
		ctx,
		uid,
		req.Email,
		int(roleID),
	)
}

// Hàm này dùng để tạo ra tài khoản người dùng mới với Transaction
func (l *LoginWithGoogleUseCase) createUserWithTransaction(
    ctx context.Context,
    req *models.CreateUsersRequestNonStrict,
) (string, error) {
    var uid string

    err := servicesupport.RunInTx(ctx,
        l.db,
        l.zapLogger,
        func(tx *sql.Tx) error {
            userRepoTx := l.userRepo.WithTx(tx)
            userRoleRepoTx := l.userRoleRepo.WithTx(tx)

            createdUID, err := userRepoTx.CreateUserFromOAuth2(ctx, req)
            if err != nil {
                l.slog.LogError(
                    "Lỗi tạo mới người dùng từ OAuth2",
                    err,
                    zap.Error(err),
                )
                return err
            }
			uid = createdUID

            if _, err := userRoleRepoTx.Create(ctx, &models.UserRole{
                Id_role: defaultUserRoleID,
                Uuid:     uid,
            }); err != nil {
				l.slog.LogError(
					"Tạo role mặc định cho user Google thất bại",
					err,
					zap.String("uuid", uid),
					zap.Int32("role_id", defaultUserRoleID),
				)
                return err
            }

            return nil
        },
    )

    return uid, err
}