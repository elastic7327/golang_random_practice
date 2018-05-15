package interfaces

import (
	"fmt"
	"strconv"
	"testing"
)

type Human struct {
	name string
	age  int
}

func formatString(arg interface{}) string {

	// 빈 인터스페이스를 하게되면 모든 타입을 받을 수가 있습니다.
	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)

	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)

	case string:
		f := arg.(string)
		return f

	case Human:
		p := arg.(Human)
		return p.name + " " + strconv.Itoa(p.age)

	case *Human:
		p := arg.(*Human)
		return p.name + " " + strconv.Itoa(p.age)

	default:
		return "Error"
	}

}

func TestEmptyInterfaces(t *testing.T) {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world"))

	fmt.Println(formatString(Human{"Maria", 20}))
	fmt.Println(formatString(&Human{"Jonni", 12}))

}

func TestEmptyInterfaceTypeCheck(t *testing.T) {

	var z interface{}

	z = Human{"Red Bull", 881004}

	if v, ok := z.(Human); ok {
		fmt.Println(v, ok)
	}
}
