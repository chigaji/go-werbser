package main

import (
	"fmt"
)

func main() {

	num3 := 3
	doubleNo := func() int {
		num3 *= 2
		// fmt.Println(num3)
		return num3

	}
	fmt.Println(doubleNo())
}
