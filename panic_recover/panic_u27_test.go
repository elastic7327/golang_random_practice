package panic

import (
	"fmt"
	"testing"
)

// recovery func
func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	panic("Error !!")

}

func TestPanic(t *testing.T) {
	f()
	fmt.Println("Hello World")
}
