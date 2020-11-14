package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return " main task Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("async returned result.")
		retCh <- ret
		fmt.Println("async service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	//ret := <- AsyncService()
	//t.Log(ret)
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
