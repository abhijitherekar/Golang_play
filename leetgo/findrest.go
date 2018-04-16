package main

import (
	"fmt"
	"math"
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
				ans = []string{str2}
				min = currmin
			}
		}
	}
	return ans
}

func finddup(nums []int) bool {
	dupmap := make(map[int]bool)
	for _, num := range nums {
		if !dupmap[num] {
			dupmap[num] = true
		} else {
			return true
		}
	}
	return false
}

func finddupdiff(nums []int, k int) bool {
	dupmap := make(map[int]int)
	for i, num := range nums {
		if !(dupmap[num] > 0) {
			dupmap[num] = i + 1
		} else {
			if int(math.Abs(float64(dupmap[num]-i-1))) <= k {
				return true
			}
		}
	}
	return false
}

func main() {
	list1 := []string{"a", "b", "c", "z"}
	list2 := []string{"z", "a", "e", "i"}
	list3 := []int{1, 2, 3, 4, 1}
	fmt.Println(findRestaurant(list1, list2))
	fmt.Println(finddupdiff(list3, 3))
}
