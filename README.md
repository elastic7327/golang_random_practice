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

vim-go의 built-in debugger을 사용하는방법..

vim-go 설치이후 빔내에서 `F9`를 눌러버리면, `>` 이런식으로 브레이크 포인터가 생긴다.
이후 사용방법은 vim-go 레퍼런스에 자세하게 나와있다. https://github.com/fatih/vim-go/blob/master/doc/vim-go.txt

`사실 vim-go 를 적절하게 쓰면 delv를 쓸 필요조차 없다`
`F9  Breakpoint`, `F8 Continue`, `F10 Next`, `F11 Step`
```
:GoDebugBreakpoint [linenr]

Toggle breakpoint for the [linenr]. [linenr] defaults to the current line
if it is omitted. A line with a breakpoint will have the
{godebugbreakpoint} |:sign| placed on it. The line the program is
currently halted on will have the {godebugcurline} sign.

            *hl-GoDebugCurrent* *hl-GoDebugBreakpoint*
                A line with a breakpoint will be highlighted with the {GoDebugBreakpoint}
                    group; the line the program is currently halted on will be highlighted
                        with {GoDebugCurrent}.

                            Mapped to <F9> by default.  
*:GoDebugContinue*
*(go-debug-continue)*


:GoDebugContinue

    Continue execution until breakpoint or program termination. It will start
        the program if it hasn't been started yet.

    Mapped to <F5> by default.

*:GoDebugNext*
*(go-debug-next)*

```
