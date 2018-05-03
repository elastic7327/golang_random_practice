package slices

import (
	"fmt"
	"testing"
)

func TestSlicesU22(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	x := append(a, b...) // 슬라이스 a에 슬라이스 b를 붙일 때는 b...를 쓴다

	fmt.Println(x)

}
