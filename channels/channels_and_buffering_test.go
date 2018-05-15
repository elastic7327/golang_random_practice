package channels

import (
	"fmt"
	"testing"
)

func TestChannelsAndBuffering(t *testing.T) {

	// 채널의 버퍼가 가득차면 값을 꺼내서 출력합니다.
	// 채널에 버퍼를 1개 이상 설정하면 비동기 채널(asynchronous channel)이 생성됩니다.
	// 비동기 채널은 보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고 대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기합니다.

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("GoRoutine: ", i)
		}
	}()

	// 고루틴을 생성하고, 반복문을 실행할 때마다 채널 done에 true 값을 보냅니다.

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("Main Func", i)
	}

	// WIP

}
