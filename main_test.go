package main

import (
	"testing"
)

func Test_SayHello(t *testing.T) {
	var testCases = map[string]string{
		"C/C++":      "Hello, C/C++.",
		"Java":       "Hello, Java.",
		"JavaScript": "Hello, JavaScript.",
		"Golang":     "Hello, Golang.",
		"PHP":        "Hello, PHP.",
		"Python":     "Hello, Python.",
		"Lua":        "Hello, Lua.",
	}

	for key, expected := range testCases {
		result := SayHello(key)
		if result != expected {
			t.Errorf("Test_SayHello failed. Input[%v], Expected[%v], but Got[%v].", key, expected, result)
		}
	}
}
