package string

import "testing"

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
