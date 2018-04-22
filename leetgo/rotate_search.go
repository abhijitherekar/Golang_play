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

//return index of the searched element
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	fmt.Println(nums)

	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		//case 1 mid == target
		if nums[mid] == target {
			return mid
		}

		fmt.Println("low: ", low, "hi", high, "nums[mid]: ", nums[mid], "mid: ", mid, "a[low]", nums[low])

		// target < a[mid], which can be on either side.
		// so how to decide which side to go???
		//case 2, if a[mid] < a[high], then we know a[mid....high]
		// is sorted. so, just call BS on a
		if nums[mid] < nums[high] {
			// target in a[mid...high]
			if target >= nums[mid] && target <= nums[high] {
				low = mid
			} else if target >= nums[low] {
				//target in a[low...mid-1]
				high = mid - 1
			}
			//	continue
		}
		//case 3, if a[mid] > a[low], then we know the a[low....mid] is sorted
		if nums[mid] > nums[low] {
			if target <= nums[mid] && target >= nums[low] {
				high = mid - 1
			} else if target <= nums[high] {
				low = mid + 1
			} else {
				return -1
			}
		}
		if high == low && nums[high] == target {
			return high
		}
	}
	return -1
}

func main() {
	//[4,5,6,7,0,1,2]
	a := []int{3, 1, 2}
	//a := []int{12, 4, 5, 6, 7, 10, 11}
	//end := len(a)
	//fmt.Println(findstart(a))
	fmt.Println(search(a, 1))
}
