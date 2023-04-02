package util

import "golang.org/x/crypto/bcrypt"

// PasswordHash Creates a password hash
// 密码加密: 同PHP函数 password_hash()
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// PasswordVerify : Verifies that a password matches a hash
// 加密密码验证: 同PHP函数 password_verify()
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
