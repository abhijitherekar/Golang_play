package main

import (
	"fmt"
	"math"
	//"strings"
	"bytes"
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

var map1 = map[int]string{
	1: "M",
	2: "MM",
	3: "MMM",
}
var map2 = map[int]string{
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}
var map3 = map[int]string{
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}
var map4 = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

func inttorom(n int) string {
	var buffer bytes.Buffer
	for i := 3; i >= 0; i-- {
		if i == 3 {
			buffer.WriteString(map1[n/int(math.Pow(10, float64(i)))])
		}
		if i == 2 {
			buffer.WriteString(map2[n/int(math.Pow(10, float64(i)))])
		}
		if i == 1 {
			buffer.WriteString(map3[n/int(math.Pow(10, float64(i)))])
		}
		if i == 0 {
			buffer.WriteString(map4[n/int(math.Pow(10, float64(i)))])
		}
		n = n % int(math.Pow(10, float64(i)))
		fmt.Println(n, buffer.String())
	}
	return "hi"
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
	fmt.Println(inttorom(3999))
}
