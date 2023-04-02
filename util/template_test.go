package util

import (
	"testing"
)

// IsAvailable : check key is_avail
func Test_IsAvailable(t *testing.T) {
	var res bool
	var expected bool

	expected = true
	res = IsAvailable("key_name", map[string]string{"key_name": "ni hao"})
	if res != expected {
		t.Errorf("Test_IsAvailable failed. Excpected[%v] Got[%v]", expected, res)
	}
	res = IsAvailable("KeyName", struct{ KeyName string }{"test struct"})
	if res != expected {
		t.Errorf("Test_IsAvailable failed. Excpected[%v] Got[%v]", expected, res)
	}

	expected = false
	res = IsAvailable("not_exist_key", map[string]string{"name": "hi failed"})
	if res != expected {
		t.Errorf("Test_IsAvailable failed. Excpected[%v] Got[%v]", expected, res)
	}
	res = IsAvailable("not_exist_key", struct{ KeyName string }{"test struct failed"})
	if res != expected {
		t.Errorf("Test_IsAvailable failed. Excpected[%v] Got[%v]", expected, res)
	}
}
