// given the 32-bit signed integer reverse it.
//example: 123 will be 321

package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(num int) int {
	maxint := 0x7fffffff
	x := math.Abs(float64(num))
	result := 0
	strlen := len(strconv.Itoa(int(x))) - 1
	for x != 0 {
		result += (int(x) % 10 * int(math.Pow(10, float64(strlen))))
		x = x / 10
		strlen -= 1
	}
	if result > maxint {
		return 0
	}
	if num < 0 {
		return result - 2*result
	}
	return result
}

func main() {
	a := -320
	fmt.Println(reverse(a))
}
