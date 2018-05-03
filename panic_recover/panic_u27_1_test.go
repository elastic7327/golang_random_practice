package panic

import (
	"fmt"
	"testing"
)

func f() {

	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}

}

func TestPanicTwo(t *testing.T) {

	f()

	fmt.Println("Hello, World")

}
