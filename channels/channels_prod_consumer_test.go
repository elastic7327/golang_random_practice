package channels

import (
	"fmt"
	"testing"
)

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100
}

func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

func awesum(a, b int) <-chan int {

	out := make(chan int)

	go func() {

		out <- a + b

	}()

	return out
}

func TestProducerAndConsumer(t *testing.T) {

	t.Skip()

	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()

}

func TestAweSumFuncTest(t *testing.T) {

	z := awesum(1, 2)

	fmt.Println(<-z) // 채널이기 때문에 채널 값을 가져오는 <- 를 사용해야합니다.

}
