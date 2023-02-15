package main

import "fmt"

func main() {
	str := "sdfihoiaxcjmofcisaerioljveghkklklrnvmaognoahxcm"

	var count [26]int // 영문자 수
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}
	fmt.Println(maxCount)
	fmt.Printf("Max charater: '%c' count: '%d'", maxCh, maxCount)
}
