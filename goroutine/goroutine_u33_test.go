package goroutine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func hello() {
	fmt.Println("Hello, world!")
}

func gollo(n int) {
	r := rand.Intn(100)

	time.Sleep(time.Duration(r))

	fmt.Println(n)
}

func TestGoroutine(t *testing.T) {

	go hello()

	fmt.Scanln() // main 함수가 종료되지 않도록 대기
}

func TestGoRoutineWithRandit(t *testing.T) {

	/*
		고루틴 100개를 동시에 실행하면서 각 고루틴은 랜덤한 시간동안 대기합니다.
		rand.Intn 함수에 100을 넣으면 0부터 100까지 랜덤한 숫자를 리턴합니다.
		그리고 time.Sleep 함수를 사용하여 설정한 시간동안 대기합니다.
		여기서 time.Sleep 함수에 값을 넣을 때는 time.Duration 타입으로 변환해줍니다.
	*/

	for i := 0; i < 100; i++ {
		go gollo(i)
	}

	fmt.Scanln()

}
