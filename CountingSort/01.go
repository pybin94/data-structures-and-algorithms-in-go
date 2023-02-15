package main

import "fmt"

func main() {

	arr := []int{5, 4, 2, 7, 3, 4, 7, 4, 2, 3, 5, 0, 8, 10, 9, 6, 4, 6, 1, 8, 5, 2, 10, 5, 6, 3, 7, 3, 0, 4, 1, 7, 4, 2}
	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	sorted := make([]int, 0, len(arr))

	// 2중 for문 중복된 숫자가 많을 수록 비 효율적
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println("count:", count)
	fmt.Println("sorted:", sorted)

}
