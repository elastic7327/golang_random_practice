// Package  mutexs provides ...
package mutexs

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexReadWrite(t *testing.T) {
	t.Skip("Skip for moment")
	var data int = 0

	go func() {

		for i := 0; i < 3; i++ {
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
		}

	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read: 1", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("read: 2", data)
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)

}

func TestMutexReadWriteWithLock(t *testing.T) {

	var data int = 0

	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Microsecond)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 1 : ", data)
			time.Sleep(10 * time.Microsecond)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(10 * time.Microsecond)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(5 * time.Second)

}
