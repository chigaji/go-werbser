package main

import "fmt"

type name struct {
	firstName string
	LastName  string
}

func main() {

	c := new(name)
	fmt.Println(c)

}
