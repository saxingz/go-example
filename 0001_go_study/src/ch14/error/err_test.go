package error

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("should not be less than 2")
var LargerThenHundredError = errors.New(" shoue be less than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 2 || n > 100 {
	//	return nil, errors.New("must in [2, 100]")
	//}

	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThenHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
