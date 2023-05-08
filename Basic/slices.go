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
	//Разворачивает слайс в обратном порядке, возвращает результат
	var h [20]int
	j := h[:0]
	for i := len(s); i > 0; i-- {
		a := i - 1
		j = append(j, s[a])
	}
	return j
}

func CreateSlice(n int) []int {
	//Создает слайс от 1 до n, возвращает результат
	s := array[:0]

	for i := 1; i < (n + 1); i++ {
		s = append(s, i)
	}

	return s
}
