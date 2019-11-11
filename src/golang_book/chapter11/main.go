package main

import (
	"fmt"
	"time"
)

// "golang_book/chapter11/math"

func main() {
	// xs := []float64{1, 2, 3, 4}
	// // avg := math.Average(xs)
	// fmt.println(avg)
	const PI = 3.1425
	fmt.Println(PI)
	fmt.Printf("type = %T\n", PI)
	num := complex(2, 5)
	fmt.Printf("Value = %v, type = %T\n", num, num)

	i := 2
	j := 3
	println("i^j = ", i^j)

	chan1 := make(chan string)
	chan1 <- "Welcome home son"
	go SendMessage(chan1)
	//fmt.Println(recieve)
}
func SendMessage(c chan string) {
	fmt.Println("getting message from channel")
	time.Sleep(4 * time.Second)
	recieved := <-c
	fmt.Println(recieved)
}
