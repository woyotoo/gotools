package util

import (
	"html/template"
	"reflect"
	"regexp"
)

func Unescaped(x string) interface{} {
	// x2 := html.EscapeString(x)
	return template.HTML(x)
}

// IsAvailable isset()
func IsAvailable(name string, data interface{}) bool {
	// 检查key值是否存在
	v := reflect.ValueOf(data)

	// fmt.Println("-------> name: ", name)
	// fmt.Println("-------> data: ", data)
	// fmt.Println("-------> v.Kind(): ", v.Kind())
	if !v.IsValid() { // 值不存在
		return false
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// 如果是map类型，使用 MapIndex
	if v.Kind() == reflect.Map {
		//fmt.Println("=======> ", v.MapIndex(reflect.ValueOf(name)))
		//fmt.Println("=======>IsValid ", v.MapIndex(reflect.ValueOf(name)).IsValid())
		return v.MapIndex(reflect.ValueOf(name)).IsValid()
	}
	// 如果是 struct 类型，使用 FieldByName
	if v.Kind() == reflect.Struct {
		return v.FieldByName(name).IsValid()
	}

	return false
}

// PureHtmlContent 获取 HTML 标签中的纯文本
func PureHtmlContent(html string) (content string) {
	reg := regexp.MustCompile(`(<[\S\s]+?>)|([\s]+?)`)
	content = reg.ReplaceAllString(html, "")
	return
}
