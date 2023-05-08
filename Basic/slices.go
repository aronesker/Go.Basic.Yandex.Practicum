package main

import "fmt"

var array [1]int

func main() {
	fmt.Println(ReverseSlice(EditSlice(10, 89)))
}

func EditSlice(n int, m int) []int {
	//Удаляет из слайса значения с индексами от n до m, возвращает результат
	t := CreateSlice(100)
	if len(t) != 0 && n < len(t) {
		t = append(t[:n], t[m+1:]...)
	}
	return t
}

func ReverseSlice(s []int) []int {
	//Разворачивает переданный слайс в обратном порядке, возвращает результат
	for i := range s[:len(s)/2] {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func CreateSlice(n int) []int {
	//Создает слайс от 1 до n, возвращает результат
	s := array[:0]

	for i := 1; i < (n + 1); i++ {
		s = append(s, i)
	}

	return s
}
