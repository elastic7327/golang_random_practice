package interfaces

import (
	"fmt"
	"testing"
)

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("Quaacckkk!")
}

func (d Duck) feathers() {
	fmt.Println("Duck has white and gray ...")
}

type Person struct {
	name string
	age  int
}

func (p Person) quack() {
	fmt.Println("(Mocking) -- Quaacckkk! #$@@#$?")
}

func (p Person) feathers() {
	fmt.Println("Person ..gray .. white..")
}

type Quacker interface {
	/* TODO: add methods */
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func TestDuckTypingExample(t *testing.T) {

	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)
}

func TestDuckTestInterfaceBool(t *testing.T) {

	var donald Duck

	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}
}
