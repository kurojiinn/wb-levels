package main

import (
	"fmt"
)

type Human struct {
}

func (h *Human) walk() {
	fmt.Println("iam walking")
}

type Action struct {
	Human
}

func main() {
	human := &Human{}

	action := &Action{}

	human.walk()
	action.walk()

}
