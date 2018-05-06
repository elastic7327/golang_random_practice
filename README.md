# golang_random_practice



`dlv` 사용법은 간단하다.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.Breakpoint() // 이런식으로 등록을 해주고..
	a := 1 // 이후에 dlv 실행후 print a 이런식으로 찍어볼 수가있음
	fmt.Println("Hello world!")
	fmt.Println(a)
}
```


```cmd

dlv debug 또는 dlv test 를 해서.. gdb 방식으로 debug 를 하면된다.
continue, print, restart 등등 유용한 명령어가 많이 있다.

(dlv) l
Stopped at: 0x588919c
=>no source available
(dlv) ls
Stopped at: 0x588919c
=>no source available
(dlv) c
> main.main() ./main.go:12 (PC: 0x10aea94)
>      7:
>      8: func main() {
       9:
       10:    runtime.Breakpoint()
       11:
=>     12:         a := 1
13:
14:          fmt.Println("Hello world!")
15:
16:          fmt.Println(a)
17: }

(dlv)
>           }
```
