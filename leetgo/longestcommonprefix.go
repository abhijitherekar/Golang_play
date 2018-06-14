package main

import (
	"bytes"
	"fmt"
)

func longestCommonPrefix(str []string) string {
	//geeksfor, geeks, gee
	var commonprefix bytes.Buffer
	prefix := str[0] //prefix = geeksfor
	i, j := 0, 0
	lenstr := len(str)
	for j < lenstr-1 {
		s1 := str[j+1]
		if (i == len(s1)) || (i == len(prefix)) {
			i = 0
			j++
			prefix = commonprefix.String()
			commonprefix.Reset()
			continue
		}
		if (prefix[i] != s1[i]) || (i == len(s1)) || (i == len(prefix)) {
			i = 0
			j++
			prefix = commonprefix.String()
			commonprefix.Reset()
			fmt.Println(i, j, "***")
			continue
		}
		//fmt.Println(s[i], s1[i], i, j)
		commonprefix.WriteByte(prefix[i])
		i++
	}
	return prefix
}

func main() {

	str := []string{"a", "b"}
	fmt.Println(longestCommonPrefix(str))
}
