package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := 1

	runtime.Breakpoint()

	fmt.Println("Hello world!")

	fmt.Println(a)

}
