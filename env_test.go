package etc

import (
	"testing"
)

func TestNewEnv(test *testing.T) {
	env := NewEnv(map[string]interface{}{"a": 1, "b": true, "name": "etc"})
	if 1 != env.Get("a") {
		test.Fail()
	}
	if true != env.Get("b") {
		test.Fail()
	}
	if "etc" != env.Get("name") {
		test.Fail()
	}
}
