package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("hey")
	getStruct()
	looping()
	timeDay()
	arrayT()
	sliceT()
	mapT()
	fmt.Println(person{"Bobby", 20})
	fmt.Println(person{name: "ken", age: 30})
	fmt.Println(&person{name: "Ann", age: 20})

	myrec := rect{with: 16, length: 8}
	fmt.Println(myrec.area())

	gochannel()

	var p = fmt.Println

	p("Golden river of Health")
}
func getStruct() {
	const message = "hello there"
	fmt.Println(message)
	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

}
func looping() {
	for {

		fmt.Println("fun")
		break

	}
}
func timeDay() {
	fmt.Println(time.Now().Weekday())
	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println(`it's after noon`)
	default:
		fmt.Println(`It's before noon`)
	}
}
func arrayT() {
	var a [5]int
	fmt.Println("emp: ", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)
}
func sliceT() {
	s := make([]string, 3)
	s[0] = "rony"
	s[1] = "mark"
	s[2] = "molly"
	s = append(s, "chigaji")

	fmt.Println("s: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)
}
func mapT() {
	m := make(map[string]string)
	m["name"] = "toyota"
	m["year"] = "2018"
	m["price"] = "280,000usd"
	fmt.Println("m:", m)

	v := m["year"]
	fmt.Println("v: ", v)
	delete(m, "price")
	fmt.Println("m: ", m)
	_, prs := m["name"]
	fmt.Println("psr :", prs)
}

type person struct {
	name string
	age  int
}
type rect struct {
	with, length int
}

func (r *rect) area() int {

	return r.length * r.with
}

func gochannel() {

	info := make(chan string)
	go func() {
		info <- "rony called"
	}()
	msg := <-info
	fmt.Println(msg)
}
