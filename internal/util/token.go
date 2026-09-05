package util

import (
	"errors"
	"strings"
	"time"
	"crypto/sha256"
	"encoding/hex"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Tạo AccessToken
func GenerateAccessToken(userID, email string, userRole int) (string , error){
	
	// Lấy secret key từ config để làm secret key cho JWT.
	secretKey := global.Config.Authentication.JWT.Secret

	token := jwt.New(jwt.SigningMethodHS256) // Tạo một token mới với phương thức ký HS256

	// Tạo claims cho token
	claims := token.Claims.(jwt.MapClaims)
	claims["jwt_id"] = uuid.NewString() // Tạo một UUID ngẫu nhiên cho jwt_id
	claims["user_id"] = userID
	claims["email"] = email
	claims["user_role"] = userRole
	claims["iss"] = global.Config.Authentication.JWT.Issuer
	claims["token_type"] = "access_token"
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(global.Config.Authentication.JWT.AccessTokenExpirationMinutes)).Unix()

	// Ký token với secret key
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Tạo refresh token, bằng cách tạo một chuỗi ngẫu nhiên có độ dài 64 ký tự
// Lưu ý: refresh token không cần phải có claims, vì nó chỉ được sử dụng để lấy access token mới
func GenerateRefreshToken() (string, error){

	// 1. Tạo chuỗi ngẫu nhiên có độ dài 64 ký tự
	rawToeken := GenerateRandomString(64)

	return rawToeken, nil
}

// ParseToken: Hàm này sẽ parse Bearer token từ header Authorization và trả về header và payload , ngược lại trả về rỗng
func ParseToken(authHeader string) (string, string) {
	const prefix = "Bearer "

	// Kiểm tra xem header có bắt đầu bằng "Bearer " hay không
	// if strings.HasPrefix(authHeader, "Bearer ") {
	// 	token := strings.TrimPrefix(authHeader, "Bearer")	// Loại bỏ "Bearer " khỏi header
	// 	parts := strings.Split(token, ".") // Tách token thành 3 phần: header, payload, signature
		
	// 	// Nếu token có 2 phần (header và payload), trả về header và payload, ngược lại trả về token và rỗng
	// 	if len(parts) == 2 {
	// 		return parts[0], parts[1]
	// 	}
	// 	return token, ""
	// }
	// return "", ""

	// Kiểm tra xem header có bắt đầu bằng "Bearer
	if !strings.HasPrefix(authHeader, prefix) {
		return authHeader, ""
	}

	// Loại bỏ "Bearer " khỏi header
	token := strings.TrimPrefix(authHeader, prefix)


	// Tách token thành 3 phần: header, payload, signature
	parts := strings.Split(token,".")
	if len(parts) == 3 {
		return parts[0], parts[1]
	}
	
	return token, ""
}

// ParseTokenWithClaims: Hàm này sẽ parse JWT token và trả về claims, ngược lại trả về rỗng
func ParseTokenWithClaims(tokenStr, secretKey string) (jwt.MapClaims, error) {
	
	// Thực hiện parse token với secret key
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		
		// Kiểm tra phương thức ký của token có phải là HMAC hay không
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(secretKey), nil
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
func GetUserIDFromClaimsInHeader(ctx *fiber.Ctx, secretKey string) (string, error)  {
	
	token := ctx.Get("Authorization") // Lấy token từ header Authorization
	if token == "" {
		return "", errors.New("Authorization header is missing")
	}

	// loại bỏ "Bearer " khỏi token
	token = strings.TrimPrefix(token, "Bearer ")

	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(token, secretKey)
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
	
	secretKey := global.Config.Authentication.JWT.Secret
	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(tokenStr, secretKey)
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
	secretKey := global.Config.Authentication.JWT.Secret
	// Parse token và lấy claims
	claims, err := ParseTokenWithClaims(tokenStr, secretKey)
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
	secretKey := global.Config.Authentication.JWT.Secret
	claims, err := ParseTokenWithClaims(tokenStr, secretKey)
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