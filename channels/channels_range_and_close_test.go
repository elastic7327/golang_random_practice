package channels

import (
	"fmt"
	"testing"
)

func TestRangeAndClose(t *testing.T) {

	c := make(chan int) // int 형 채널을 생성

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄
		}
		close(c) // 채널을 닫음
	}()
	/*
		for 반복문 안에서 range 키워드를 사용하면 채널이 닫힐 때까지 반복하면서 값을 꺼냅니다.
		여기서는 동시에 고루틴 안에서 채널 c에 0부터 4까지 값을 보낸 뒤 close를 사용하여 채널을 닫습니다.
		이렇게 하면 range로 0부터 4까지 꺼내고, 값을 출력한 뒤 반복문이 종료됩니다.
	*/
	for v := range c {
		fmt.Println(v)

	}

	/*
		이미 닫힌 채널에 값을 보내면 패닉이 발생합니다.
		채널을 닫으면 range 루프가 종료됩니다.
		채널이 열려있고, 값이 들어오지 않는다면
		range는 실행되지 않고 계속 대기합니다.

		만약 다른 곳에서 채널에 값을 보냈다면(채널에 값이 들어오면) 그때부터 range가 계속 반복됩니다.

	*/

	k := make(chan int, 1)

	go func() {
		k <- 1
	}()

	a, ok := <-k

	fmt.Println(a, ok)

	close(k)

	a, ok = <-k
	fmt.Println(a, ok)

}

func TestRangeWorks(t *testing.T) {

	// 위쪽 테스트 에서 혼란이 올 수도 있는데 단순 5에 대해서 range를 할때에 Python 문법처럼 혼란이 오면 안됩니다.
	// 버퍼에 넣는다 라고 생각하면 편합니다. 아래 출력결과를 확인하면 보다 더 쉽게 이해할수 있습니다.

	cc := make(chan int)

	go func() {
		cc <- 12
		cc <- 3
		cc <- 2
		cc <- 100
		close(cc)
	}()

	for v := range cc {
		fmt.Println(v)
	}
}
