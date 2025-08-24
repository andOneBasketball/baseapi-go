package xcrypto

import (
	"testing"
)

func TestEncryptAESBase64(t *testing.T) {
	key := []byte("12345678901234567890123456789012") // 32 bytes
	plaintext := "hello world"

	encrypted, err := EncryptAESBase64(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptAESBase64 failed: %v", err)
	}

	if encrypted == "" {
		t.Error("Encrypted result should not be empty")
	}

	// 测试解密
	decrypted, err := DecryptAESBase64(encrypted, key)
	if err != nil {
		t.Fatalf("DecryptAESBase64 failed: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decrypted result '%s' does not match original '%s'", decrypted, plaintext)
	}
}

func TestDecryptAESBase64_InvalidKey(t *testing.T) {
	key := []byte("12345678901234567890123456789012")
	wrongKey := []byte("09876543210987654321098765432109")
	plaintext := "hello world"

	encrypted, err := EncryptAESBase64(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptAESBase64 failed: %v", err)
	}

	// 使用错误的密钥解密
	_, err = DecryptAESBase64(encrypted, wrongKey)
	if err == nil {
		t.Error("Expected error when using wrong key")
	}
}

func TestDecryptAESBase64_InvalidInput(t *testing.T) {
	key := []byte("12345678901234567890123456789012")

	// 测试无效的 base64 字符串
	_, err := DecryptAESBase64("invalid-base64", key)
	if err == nil {
		t.Error("Expected error for invalid base64 input")
	}

	// 测试太短的密文
	_, err = DecryptAESBase64("dGVzdA==", key) // "test" in base64
	if err == nil {
		t.Error("Expected error for too short ciphertext")
	}
}
