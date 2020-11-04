package ch13

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.println"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v \n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	writeFirstProgram(goProg)
	goProg2 := &GoProgrammer{}
	writeFirstProgram(goProg2)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(javaProg)
}
