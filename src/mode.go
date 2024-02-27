package main

import "fmt"

func main() {

	// numbers := [5]int{1,2,3,4,5}
	arr := []int{1, 1, 2, 3, 3, 3, 4, 4, 5, 1000}
	fmt.Println(arr)
	mode := modeCalc(arr)
	fmt.Printf("Mode: %d", mode)
}

func modeCalc(arr []int) int {
	m := map[int]int{}
	var mfreq, maxf int // mfreq - кол-во одинаковых чисел
	if len(arr) == 1 {
		return arr[0]
	}
	for i, a := range arr {
		m[a]++
		if i == 0 { // если i первое число запиши в maxf его ключ и в mfreg его кол-во повторений
			maxf = a
			mfreq = 1
		} else { // иначе
			if m[a] > mfreq { // если кол-во ключей больше чем в с предыдушим число
				maxf = a     // то запиши это ключ в maxf
				mfreq = m[a] // а в mfreq его кол-во одинаковых чисел
			} else if m[a] == mfreq && a < maxf { // иначе если у нас есть одинаковое кол-во повторений чисел, то запиши ключъ который меньше
				maxf = a
			}
		}
	}
	return maxf
}
