package pointer

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	var numPtr = new(int)

	fmt.Println(numPtr)

	*numPtr = 1

	fmt.Println(*numPtr)

}

func hello(n *int) {
	*n = 2
}

func TestPointerFunc(t *testing.T) {

	n := 1
	hello(&n)
	fmt.Println(n)
}
