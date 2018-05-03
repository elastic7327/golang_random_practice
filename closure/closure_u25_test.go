package closure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosure(t *testing.T) {

	assert.Equal(t, 1, 1)

	sum := func(a, b int) int { return a + b }

	r := sum(1, 2)

	fmt.Println(r)

	a, b := 3, 5

	f := func(x int) int {

		return a*x + b // 함수 바깥의 변수 a, b 사용

	}

	y := f(5)

	fmt.Println(y)

}

func calc() func(x int) int {
	a, b := 3, 5

	return func(x int) int {
		return a*x + b
	}

}

func TestClosureFunc(t *testing.T) {

	f := calc()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))

}
