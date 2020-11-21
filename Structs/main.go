package main

import (
	"fmt"
)

type Person struct {
	Name string
	Id   int
}

func main() {
	person1 := Person{Name: "sakib", Id: 1}
	fmt.Println(person1)
	fmt.Println(person1.Name, person1.Id)
}
