package loop_test

import "testing"

/**
@author saxing 2020/9/28 22:21
*/
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
