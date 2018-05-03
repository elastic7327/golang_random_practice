package method

import (
	"fmt"
	"testing"
)

type Rectangle struct {
	width, height int
}

func (rect *Rectangle) scaleA(factor int) {

	rect.width = rect.width * factor
	rect.height = rect.height * factor

}

func (rect Rectangle) scaleB(factor int) {

	rect.width = rect.width * factor
	rect.height = rect.height * factor

}

func TestMethods(t *testing.T) {

	r := Rectangle{30, 30}

	r.scaleA(2)

	fmt.Println(r)

	r.scaleB(5)

	fmt.Println(r)

}
