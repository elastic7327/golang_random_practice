package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMultiCoreGoroutine(t *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello Daddy"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()

}
