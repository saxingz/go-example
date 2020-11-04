package encap2

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Logf(e1.String())
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

/**
通过实例访问
*/
func (e Employee) String() string {
	fmt.Printf(" -- 2 Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

/**
通过指针访问
推荐写法，因为不会产生数据复制
*/
func (e *Employee) String2() string {
	fmt.Printf(" -- 3 Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s--Name:%s--Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf(" -- 1 Address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	e2 := &Employee{"0", "Bob", 20}
	fmt.Printf(" -- 1 Address is %x \n", unsafe.Pointer(&e2.Name))
	t.Log(e2.String())
	e3 := Employee{"0", "Bob", 20}
	fmt.Printf(" -- 1 Address is %x \n", unsafe.Pointer(&e3.Name))
	t.Log(e3.String2())
	e4 := &Employee{"0", "Bob", 20}
	fmt.Printf(" -- 1 Address is %x \n", unsafe.Pointer(&e4.Name))
	t.Log(e4.String2())
}
