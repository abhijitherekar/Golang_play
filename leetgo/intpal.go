package main

import (
	"fmt"
	//"github.com/Golang_play/leetgo/reverseint"
	"math"
)

func ispalindrome(x int) bool {
	if x < 0 {
		return false
	}
	/* using the reverse int we can do this
	if x == reverseint.Reverse(x) {
		return true
	}
	return false
	*/
	num := x
	//the 2nd method that we will try to do the bit wise way
	// here we will counts number of digits
	digits := 0
	for {
		if x < 10 {
			break
		}
		x = x / 10
		digits++
	}

	i := 0
	// Now here we need to walk digit by digit
	//1st we will start from left most and the right most
	// in order to get the left most, we must divide by the 10^digits, num/10^digits %10
	// and to get the right most we need to divide by 10^0 %10 = num/10^0 % 10
	// and this we have to do iteratively so we need keep track of the
	// indexes
	for i < digits {
		ldig := (num / int(math.Pow(10, float64(digits)))) % 10
		rdig := (num / int(math.Pow(10, float64(i)))) % 10
		fmt.Println(ldig, rdig, digits)
		if ldig != rdig {
			return false
		}
		i++
		digits--
	}
	return true
}

func main() {
	fmt.Println(ispalindrome(1234))
	fmt.Println(ispalindrome(4334))

}
