package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	fmt.Println("Start")

	os.Exit(-1)
}

func TestPanicVxExit2(t *testing.T) {
	defer func() {
		fmt.Println("finally!")
	}()

	fmt.Println("Start")

	panic(errors.New("Something wrong!"))
}

func TestPanicVxExit3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")

	panic(errors.New("Something wrong!"))
}