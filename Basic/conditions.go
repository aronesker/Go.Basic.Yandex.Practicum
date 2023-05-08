package main

import "fmt"

var year int = 2013

func main() {
	switch {
	case year < 1945:
		fmt.Println("Слишком old!")
	case year > 1945 && year < 1965:
		fmt.Println("Привет, бумер")
	case year > 1964 && year < 1981:
		fmt.Println("Привет, представитель X!")
	case year > 1980 && year < 1997:
		fmt.Println("Привет, миллениал!")
	case year > 1996 && year < 2013:
		fmt.Println("Привет, зумер!")
	default:
		fmt.Println("Привет, альфа!")
	}
}
