package util

import (
	"errors"
	"net/url"
	"strings"
)

// 验证跳转域名的合法性：即跳转链接的域名 应该包含 baseURI 的域名
// ValidateBaseURI validates that redirectURI is contained in baseURI
func ValidateBaseURI(baseURI string, redirectURI string) error {
	base, err := url.Parse(baseURI)
	if err != nil {
		return err
	}

	redirect, err := url.Parse(redirectURI)
	if err != nil {
		return err
	}

	// log.Println("---> base.Host => ", base.Host)
	// log.Println("---> redirect.Host => ", redirect.Host)
	if !strings.HasSuffix(redirect.Host, base.Host) {
		return errors.New("invalid redirect uri")
	}

	return nil
}
