package main

import "fmt"

func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Opps", Height: 184.3},
		{Name: "Viin", Height: 163.4},
		{Name: "Tom", Height: 170.6},
		{Name: "Sim", Height: 191.5},
		{Name: "Jack", Height: 173.4},
		{Name: "Hong", Height: 169.7},
		{Name: "Gildong", Height: 188.4},
	}

	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name", name, "hieght", float64(i)/10)
		}
	}

	// 170 ~ 180
	for i := 1700; i < 1800; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name", name, "hieght", float64(i)/10)
		}
	}
}
