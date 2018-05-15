package channels

import (
	"fmt"
	"testing"
	"time"
)

func TestSyncChannels(t *testing.T) {

	done := make(chan bool)

	count := 3

	go func() {

		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("GoRoutine: ", 1) // 고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
			time.Sleep(1 * time.Second)
		}

	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main test: ", i) // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
	}

	/*
		동기 채널은 보내는 쪽에서는 값을 받을 때까지 대기하고,
		받는 쪽에서는 채널에 값이 들어올 때까지 대기합니다. 따라서, 동기 채널을 활용하면 고루틴의 코드 실행 순서를 제어할 수 있습니다.
	*/
}
