package testing

/**
test
@author saxing 2020/11/15 13:21
*/
import "testing"

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d the actual %d", inputs[i], expected[i], ret)
		}
	}
}