package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wendnesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstanTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry2(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
