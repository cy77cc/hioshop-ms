package cryptx

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// 加密
func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
