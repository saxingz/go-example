package main

import (
	"fmt"
	"os"
)

/**
main test

@author saxing 2020/7/19 19:59
*/
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}
