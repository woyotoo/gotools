package util

import (
	"testing"
)

// cd gotools
// go test util/validate_test.go util/validate.go
// go test -v util/validate_test.go util/validate.go

// TEST: ValidateBaseURI
func Test_ValidateBaseURI(t *testing.T) {
	baseURI := "https://wang123.net"
	redirectURI := "https://wang123.net/oauth2/callback"
	redirectURI_fail := "https://baidu.com/oauth2/callback"

	// 正常案例
	err := ValidateBaseURI(baseURI, redirectURI)
	if err != nil {
		t.Error("Test_ValidateBaseURI failed. Got: ", err.Error())
	}

	// 失败案例
	err = ValidateBaseURI(baseURI, redirectURI_fail)
	if err == nil {
		t.Error("Test_ValidateBaseURI failed. ", "err.Error()")
	}
	if err != nil && err.Error() != "invalid redirect uri" {
		t.Error("Test_ValidateBaseURI failed. Got: ", err.Error())
	}
}
