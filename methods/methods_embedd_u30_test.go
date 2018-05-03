package method

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {

	fmt.Println("Hello world")

}

type Student struct {
	p      Person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계
	school string
	grade  int
}

type StrangeStudent struct {
	Person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	school string
	grade  int
}

func TestMethodEmbedd(t *testing.T) {

	var s Student
	s.p.greeting()
	fmt.Println(s.p.name)
	var ss StrangeStudent
	ss.Person.greeting()

}
