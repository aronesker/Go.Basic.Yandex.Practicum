package main

import "fmt"

const (
	one = (iota + 0.5) * 2 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
