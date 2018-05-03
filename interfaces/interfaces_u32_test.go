package interfaces

import (
	"fmt"
	"testing"
)

type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	/* TODO: add methods */
	Print()
}

func TestInterfaces(t *testing.T) {

	var i MyInt = 5

	r := Rectangle{10, 20}

	var p Printer

	p = i
	p.Print()

	p = r
	p.Print()

}

func TestInterfaceArray(t *testing.T) {

	var i MyInt = 5
	r := Rectangle{10, 20}

	pArr := []Printer{i, r}

	for k := range pArr {
		pArr[k].Print()
	}

	for _, v := range pArr {
		v.Print()
	}

}
