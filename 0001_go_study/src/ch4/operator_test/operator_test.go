package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

/**

operator test

@author saxing 2020/7/26 23:20
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 6, 3, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	t.Log(3 << 0)

	a := 7 // 0111

	a = a &^ Readable
	a = a &^ Executable

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
