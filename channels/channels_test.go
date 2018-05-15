package channels

import (
	"fmt"
	"testing"
)

func sum(a int, b int, c chan int) {
	c <- a + b
}

func TestWhatIsChannels(t *testing.T) {

	// 채널을 사용하기 전에는 반드시 make 함수로 공간을 할당해야 합니다.
	//	그리고 이렇게 생성하면 동기 채널(synchronous channel)이 생성됩니다.

	c := make(chan int)

	go sum(1, 2, c)

	n := <-c

	fmt.Printf("What Happend? %d", n)

	var ch chan int

	ch = make(chan int)

	go sum(5, 6, ch)

	/*
		채널 <-: 채널에 값을 보냅니다.
		<- 채널: 채널에 값이 들어올 때까지 기다린 뒤 값이 들어오면 값을 가져옵니다.
		가져온 값은 :=, =를 사용하여 변수에 대입할 수 있습니다.
		값을 가져오면 다음 코드를 실행합니다.
	*/

	nn := <-ch

	fmt.Println(nn)
}
