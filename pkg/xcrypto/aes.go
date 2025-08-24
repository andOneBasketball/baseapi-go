package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// 生成 32 字节随机密钥（256-bit）
func GenerateRandomKey() ([]byte, string, error) {
	key := make([]byte, 32) // 256-bit
	_, err := rand.Read(key)
	if err != nil {
		return nil, "", err
	}
	return key, base64.StdEncoding.EncodeToString(key), nil
}

// 使用 AES-256-CBC 加密并 base64 编码输出
func EncryptAESBase64(plaintext string, key []byte) (string, error) {
	// 初始化 AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 使用 PKCS7 填充
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	padded := append([]byte(plaintext), padtext...)

	// 随机生成 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用 CBC 模式加密
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	// 连接 IV + ciphertext，再 base64 编码
	final := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(final), nil
}

// 解密 base64 编码的 AES 加密字符串
func DecryptAESBase64(encoded string, key []byte) (string, error) {
	// base64 解码
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	if len(raw) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// 提取 IV 和 ciphertext
	iv := raw[:aes.BlockSize]
	ciphertext := raw[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除 PKCS7 填充
	padding := int(plaintext[len(plaintext)-1])
	if padding < 1 || padding > aes.BlockSize {
		return "", fmt.Errorf("invalid padding")
	}
	return string(plaintext[:len(plaintext)-padding]), nil
}
