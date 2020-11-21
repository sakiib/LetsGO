package main

import (
	"fmt"
)

type Person struct {
	name string
	id   int
}

func main() {
	person1 := Person{name: "sakib", id: 1}
	fmt.Println(person1)
	fmt.Println(person1.name, person1.id)
}
