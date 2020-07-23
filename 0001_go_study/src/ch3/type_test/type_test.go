package type_test

import (
	"testing"
)

type MyInt int64

func TestImplict(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// go的指针无法进行运算
	//aPtr += 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}