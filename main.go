package main

import "fmt"

type People struct {
}

func (receiver People) Worlk() {
	fmt.Println("people  worlk")
}

type Man struct {
	People
}

func main() {
	fmt.Println("hello  world")
	m := Man{}
	m.Worlk()
}
