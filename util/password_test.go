package util

import (
	"testing"
)

// cd gotools
// go test -v util/password_test.go util/password.go

// Test_PasswordHash: 测试加密
func Test_PasswordHash(t *testing.T) {
	t.Log("--> Test_PasswordHash ")

	password := "admin123"
	hash, _ := PasswordHash(password)

	t.Log("--> 输入密码:", password)
	t.Log("--> 生成hash:", hash)
	// $2a$10$lRewHvjtPrYZK4TQy.TWDemBMqwIEy/.IVoB7x/2KfqrjzZJNP2ia
	// $2a$10$KEl9ZHfD4gAPu/hgXAjAm.TLgWi5ce7EzBgTdhBfW5IOimtOSfin2

	match := PasswordVerify(password, hash)
	t.Log("--> 验证结果:", match)
}

// Test_PasswordVerify: 密码验证
func Test_PasswordVerify(t *testing.T) {
	t.Log("--> Test_PasswordVerify ")

	password := "admin2"
	// PHP 生成的密码为 $2y$ 开头 (PHP实现 $2a$ 时有bug，修复时改成了 $2y$)
	hash := "$2y$10$f7ZKW1ZOR4UzGM37.GTmTuY6RmJHknfSwhBacki.Yro1I1kIddEiu"

	match := PasswordVerify(password, hash)
	t.Log("--> Test_PasswordVerify 验证结果:", match)
	if match == false {
		t.Errorf("Test_PasswordVerify failed. Got false, expected true.")
	}
}
