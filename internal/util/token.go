package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	UserRole int    `json:"user_role"`
}

type RefreshToken struct {
	UserID   string `json:"user_id"`
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
	Revoked  bool   `json:"revoked"`
}

// Tạo AccessToken
func GenerateAccessToken(userID, email string, userRole int) (string, error) {

	//1. Tạo payload cho token
	payload := &Payload{
		UserID:   userID,
		Email:    email,
		UserRole: userRole,
	}

	rawData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	//Lấy secret key & encrypt key từ config để làm secret key cho JWT.
	secretKey := global.Config.Authentication.JWT.Secret
	encryptKey := global.Config.Authentication.JWT.EncrypKey

	//3. Thực hiện mã hóa dữ liệu payload bằng AES với encryptKey từ config
	encryptedPayload, err := EncryptAES(rawData, []byte(encryptKey))
	if err != nil {
		return "", err
	}

	// 4. Tạo Claims cho token, bao gồm các thông tin cần thiết như user_id, email, user_role, issuer, token_type, iat, exp
	claims := jwt.MapClaims{
		"jwt_id":     uuid.NewString(), // Tạo một UUID ngẫu nhiên cho jwt_id
		"data":       encryptedPayload, // Lưu payload đã được mã hóa vào claims
		"iss":        global.Config.Authentication.JWT.Issuer,
		"token_type": "access_token",
		"iat":        jwt.NewNumericDate(time.Now()),
		"exp":        time.Now().Add(time.Minute * time.Duration(global.Config.Authentication.JWT.AccessTokenExpirationMinutes)).Unix(),
	}

	// 5. Tạo token mới với phương thức ký HS256 và claims đã tạo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Tạo một token mới với phương thức ký HS256

	//6. Ký token với secret key
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Tạo refresh token, bằng cách tạo một chuỗi ngẫu nhiên có độ dài 64 ký tự
// Lưu ý: refresh token không cần phải có claims, vì nó chỉ được sử dụng để lấy access token mới
func GenerateRefreshToken(uuid string) (RefreshToken, error) {

	// 1. Tạo chuỗi ngẫu nhiên có độ dài 32 ký tự
	rawToken, err := GenerateRandom(32)
	if err != nil {
		return RefreshToken{}, err
	}

	//2. Mã hóa chuỗi token vừa tạo
	token := HashToken(string(rawToken))

	return RefreshToken{
		UserID:   uuid,
		Token:    token,
		ExpireAt: time.Now().Add(time.Hour * 24 * time.Duration(global.Config.Authentication.JWT.RefreshTokenExpirationDays)).Unix(),
		Revoked:  false,
	}, nil
}

// ParseToken: Hàm này sẽ phân giải Bearer token từ header Authorization và trả về header và payload , ngược lại trả về rỗng
func ParseToken(authHeader string) (*jwt.Token, jwt.MapClaims, error) {

	// Kiểm tra xem token có mã hóa đúng thuật toán trước đó không, và kiểm tra xem token đó có đúng với chữ ký đã ký không
	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return []byte(global.Config.Authentication.JWT.Secret), nil
	})

	// Nếu có lỗi hoặc token không hợp lệ, trả về rỗng
	if err != nil || !token.Valid {
		return nil, nil, apperrors.NewInvalidTokenError("Error: Token không hợp lệ hoặc đã hết hạn")
	}

	// Lấy claims từ token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, apperrors.NewInvalidTokenError("Error: Invalid token claims")
	}

	return token, claims, nil
}

// ParseTokenWithClaims: Hàm này sẽ parse JWT token và trả về claims, ngược lại trả về rỗng
func ParseTokenWithClaims(tokenStr string) (jwt.MapClaims, error) {

	// Thực hiện parse token với secret key
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// Kiểm tra phương thức ký của token có phải là HMAC hay không
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(global.Config.Authentication.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Nếu token hợp lệ, trả về claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

// Lấy UserID từ chuỗi JWT token trong header Authorization
func GetUserIDFromClaimsInHeader(ctx *fiber.Ctx, secretKey string) (string, error) {

	token := ctx.Get("Authorization") // Lấy token từ header Authorization
	if token == "" {
		return "", errors.New("Authorization header is missing")
	}

	// loại bỏ "Bearer " khỏi token
	token = strings.TrimPrefix(token, "Bearer ")

	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(token)
	if err != nil {
		return "", err
	}

	// Lấy user_id từ claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token claims")
	}

	return userID, nil
}

// Lấy UserID từ chuỗi JWT token
func GetUserIDFromClaims(tokenStr string) (string, error) {

	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(tokenStr)
	if err != nil {
		return "", err
	}

	// Lấy user_id từ claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token claims")
	}

	return userID, nil
}

// Lấy JTI từ chuỗi JWT token
func GetJTIFromClaims(tokenStr string) (string, error) {

	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(tokenStr)
	if err != nil {
		return "", err
	}

	// Lấy jwt_id từ claims
	jwt_id, ok := claims["jwt_id"].(string)
	if !ok {
		return "", errors.New("jwt_id not found in token claims")
	}

	return jwt_id, nil
}

// Lấy thời gian hết hạn của token từ chuỗi JWT token
func GetTokenExpirationFromClaims(tokenStr string) (int64, error) {

	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(tokenStr)
	if err != nil {
		return 0, err
	}

	// Lấy exp từ claims
	expiration, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("expiration time not found in token claims")
	}

	return int64(expiration), nil
}

// hàm này dùng để băm Token, băm một cách cố định thông qua SHA256
func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// Hàm dùng để giải mã payload đã được mã hóa bằng AES từ chuỗi JWT token
func DecryptPayloadFromToken(tokenStr string) (*Payload, error) {

	// Parse token và lấy claims
	_, claims, err := ParseToken(tokenStr)
	if err != nil {
		return nil, apperrors.NewInvalidTokenError("Error: Token không hợp lệ hoặc đã hết hạn")
	}

	// Lấy cụm "data" trong claims
	encryptedData, ok := claims["data"].(string)
	if !ok {
		return nil, apperrors.NewInvalidTokenError("Error: Encode data not found in token claims")
	}

	// Giải mã dữ liệu payload bằng AES với encryptKey từ config

	decryptedBytes, err := DecryptAES(encryptedData, []byte(global.Config.Authentication.JWT.EncrypKey))
	if err != nil {
		return nil, apperrors.NewInvalidTokenError("Error: Decrypt data failed")
	}

	// Unmarshal dữ liệu payload đã giải mã thành struct Payload
	var payload Payload
	if err := json.Unmarshal(decryptedBytes, &payload); err != nil {
		return nil, apperrors.NewInvalidTokenError("Error: Unmarshal data failed")
	}

	return &payload, nil
}
