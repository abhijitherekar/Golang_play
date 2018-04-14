package main

import (
	"bytes"
	"fmt"
)

func convertToTitle(n int) string {
	var buffer bytes.Buffer
	for n > 0 {
		if n < 26 {
			buffer.WriteString(string(rune(n) + '@'))
			return buffer.String()
		}
		buffer.WriteString(string(rune(n/26) + '@'))
		fmt.Println(n)
		n = n % 26
		fmt.Println(n)
	}
	return buffer.String()
}

func main() {
	fmt.Println(convertToTitle(28))
}
