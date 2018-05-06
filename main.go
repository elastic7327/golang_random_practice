package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.Breakpoint()

	a := 1

	fmt.Println("Hello world!")

	fmt.Println(a)
}
