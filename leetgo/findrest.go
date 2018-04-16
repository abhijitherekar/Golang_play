package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	restmap := make(map[string]int)
	var ans []string
	min, currmin := 3000, 0
	for i, str := range list1 {
		restmap[str] = i + 1
	}
	fmt.Println(restmap)
	for i, str2 := range list2 {
		fmt.Println(str2, restmap[str2])
		if restmap[str2] > 0 {
			currmin = i + (restmap[str2] - 1)
			if currmin < min {
				ans = append(ans, str2)
				min = currmin
			}
		}
	}
	return ans
}

func main() {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"a", "d", "e", "b"}

	fmt.Println(findRestaurant(list1, list2))
}
