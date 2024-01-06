package foo

import "testing"

func TestFoo(t *testing.T) {
	t.Logf("TestFoo: %s", Name())
}
