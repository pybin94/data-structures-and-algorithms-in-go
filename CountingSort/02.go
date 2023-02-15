package main

import "fmt"

func main() {
	arr := []int{5, 4, 2, 7, 3, 4, 7, 4, 2, 3, 5, 0, 8, 10, 9, 6, 4, 6, 1, 8, 5, 2, 10, 5, 6, 3, 7, 3, 0, 4, 1, 7, 4, 2}

	// 중복된 숫자가 많으면 효율적
	var count2 [11]int

	for i := 0; i < len(arr); i++ {
		count2[arr[i]]++
	}

	for i := 1; i < 11; i++ {
		count2[i] += count2[i-1]
	}

	sorted2 := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted2[count2[arr[i]]-1] = arr[i]
		count2[arr[i]]--
	}

	fmt.Println("count2:", count2)
	fmt.Println("sorted2:", sorted2)
}
