package string_test

import (
	"strings"
	"testing"
)

/**
@author saxing 2020/9/30 21:20
*/
func TestString(t *testing.T) {
	var s string
	t.Log(s)
	//s = "\xE4\xBA\xBB\xFF"
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "hello"
	t.Log(s)
	//s[1] = 3

	s = "中"
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// splits
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	// join
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {

}
