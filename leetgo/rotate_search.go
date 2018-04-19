package main

import (
	"fmt"
)

// Now we know its sorted array so we need to do a Binary search
func findstart(a []int) int {
	fmt.Println(a)
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		return a[0]
	}

	mid := len(a) / 2 //7/2 = mid = 3, a[3] = 7
	start := 0
	end := len(a) - 1

	if a[mid-1] > a[mid] && a[mid] < a[(mid+1)%len(a)] {
		return a[mid]
	} else if a[mid] > a[start] && a[mid] > a[end] {
		return findstart(a[mid:])
	} else if a[mid] < a[start] && a[mid] < a[end] {
		return findstart(a[start:mid])
	}

	if (a[mid] > a[start]) && mid > start {
		return findstart(a[start:mid])
	}

	if a[mid] < a[start] && mid < end {
		return findstart(a[mid:])
	}
	return -1

}

func binarysearch(a []int, low int, high int, key int) int {

	if len(a) == 0 {
		return -1
	}
	if high < low {
		return -1
	}
	mid := low + high/2

	if a[mid] == key {
		return mid
	}
	fmt.Println(a, low, mid, high, a[mid], key)
	if a[mid] > key {
		return binarysearch(a, low, mid, key)
	} else {
		return binarysearch(a, mid, high, key)
	}
	return -1
}

//return index of the searched element
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	fmt.Println(nums)

	low := 0
	high := len(nums) - 1

	for low < high {
		mid := low + high/2
		//case 1 mid == target
		if nums[mid] == target {
			return mid
		}

		fmt.Println("target: ", target, "nums[mid]: ", nums[mid], "mid: ", mid, "a[low]", nums[low])

		// target < a[mid], which can be on either side.
		// so how to decide which side to go???
		//case 2, if a[mid] < a[high], then we know a[mid....high]
		// is sorted. so, just call BS on a
		if nums[mid] < nums[high] {
			// target in a[mid...high]
			if target >= nums[mid] {
				low = mid
			} else {
				//target in a[low...mid-1]
				high = mid - 1
			}
		}
		//case 3, if a[mid] > a[low], then we know the a[low....mid] is sorted
		if nums[mid] > nums[low] {
			if target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid
			}
		}
	}
	return -1
}

func main() {
	a := []int{11, 3, 6, 7, 8, 9}
	//a := []int{12, 4, 5, 6, 7, 10, 11}
	//end := len(a)
	//fmt.Println(findstart(a))
	fmt.Println(search(a, 6))
}
