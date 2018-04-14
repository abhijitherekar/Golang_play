package main

import (
	"fmt"
)

// Now here we will introduce the concept of rune.
// rune is nothing but a int32 or a asiic representation of the character

var mapri = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romtoint(s string) int {
	ints := []byte(s)
	fmt.Println(ints)
	runes := []rune(s)
	if len(s) == 1 {
		return mapri[runes[0]]
	}
	/*
		fmt.Println(runes)
		for i, str := range s {
			fmt.Println(i, mapri[str])
		}
	*/
	sum := 0
	ovr := false
	for i, _ := range runes {
		if ovr {
			ovr = false
			continue
		}
		if i == len(s)-1 {
			sum += mapri[runes[i]]
			return sum
		}

		if mapri[runes[i+1]] > mapri[runes[i]] {
			sum += (mapri[runes[i+1]] - mapri[runes[i]])
			ovr = true
		} else {
			sum += mapri[runes[i]]
		}
	}
	return sum
}

func main() {

	fmt.Println(romtoint("MMMCMXCIX"))

}
