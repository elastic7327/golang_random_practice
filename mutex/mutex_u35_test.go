package mutexs

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")

	var data = []int{}

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched() // 다른 고루틴이. cpu 를 사용 할 수 있도록 양보합니다.
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))

}
