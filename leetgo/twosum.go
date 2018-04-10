package main

import (
	"fmt"
	//"sort"
)

func twoSum(nums []int, target int) []int {
	if target == 0 {
		return nil
	}
	var result []int
	//var result map[int]int
	//result := make(map[int]int)
	//for now go with brute force method
	/*
		for i, _ := range nums {
			for j, _ := range nums {
				if nums[i]+nums[j] == target {
					result = append(result, i)
				}
			}
		}
	*/
	//Now for optimization, sort the slice and the use index
	/*
		sort.Ints(nums)
		start := 0
		end := len(nums) - 1

		for range nums {
			fmt.Println("\n start:", start, "end", end)
			if nums[start]+nums[end] == target {
				result = append(result, start, end)
				break
			} else if nums[start]+nums[end] > target {
				end--
			} else {
				start++
			}
		}
	*/
	//Now use the 3rd technique, if you dont have sort package, use hash table
	temp := make(map[int]int)
	for _, val := range nums {
		fmt.Println(val, "temp[val]:", temp[val])
		if temp[val] == 1 {
			result = append(result, target-val, val)
			return result
		} else {
			temp[target-val] = 1
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	a := twoSum(nums, target)
	fmt.Println(a)
}
