package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 2; i++ {
		go func() { fmt.Println(i) }()
	}
	time.Sleep(time.Second)
}
