package main

import "fmt"

type Employee interface {
	printName(name string)
	getSalary(basic float32, tax float32) float32
}
type emp string

func (e emp) printName(name string) {
	fmt.Println("employee code :", e)
	fmt.Println("employee name:", name)
}
func (e emp) getSalary(basic float32, tax float32) float32 {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var el Employee
	el = emp("001")
	el.printName("Mark")
	salary := el.getSalary(28000, 18.5)
	fmt.Println("Your salary is :", salary)
}
