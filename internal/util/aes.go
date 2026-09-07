package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

//Mã hóa và giải mã AES
func EncryptAES(plainText, key []byte) (string, error) {

	// Khởi tạo - 2 Bước đầu này luôn luôn phải có ở mã hóa và giải mã dữ liệu
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Ciphertext = nonce + ciphertext + tag
	//Dùng đến Seal để mã hóa dữ liệu, nó sẽ trả về dữ liệu đã được mã hóa
	Ciphertext := aesGCM.Seal(nonce, nonce, plainText, nil)

	return base64.URLEncoding.EncodeToString(Ciphertext), nil
}


// Giải mã AES
func DecryptAES(cipherBase64 string, key []byte) ([]byte, error) {
	// Khởi tạo - 2 Bước đầu này luôn luôn phải có ở mã hóa và giải mã dữ liệu
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Chuyển đổi dữ liệu từ base64 sang byte
	cipherText, err := base64.URLEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	// Tách nonce và ciphertext từ dữ liệu đã giải mã
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	// Dùng đến Open để giải mã dữ liệu, nó sẽ trả về dữ liệu đã được giải mã
	return aesGCM.Open(nil, nonce, cipherText, nil)
}