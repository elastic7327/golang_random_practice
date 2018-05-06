package interfaces

import (
	"fmt"
	"strconv"
	"testing"
)

func formatString(arg interface{}) {

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
		return s

	default:
		return "Error"
	}

}

func TestEmptyInterfaces(t *testing.T) {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world"))
}
