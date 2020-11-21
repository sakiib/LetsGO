package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Id   int
}

func main() {
	person1 := Person{Name: "sakib", Id: 1}
	fmt.Println(person1)
	fmt.Println(person1.Name, person1.Id)
	
	var personList []Person
	personList = append(personList, Person{Name: "messi", Id: 1})
    	personList = append(personList, Person{Name: "cr7", Id: 3})
   	personList = append(personList, Person{Name: "sakib", Id: 4})
    	personList = append(personList, Person{Name: "alamin", Id: 2})
	
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	
	// Sort by Id, keeping original order or equal elements.
    	sort.SliceStable(personList, func(i, j int) bool {
       		return personList[i].Id < personList[j].Id
    	})
	
    	for _, person := range personList {
		fmt.Println(person.Name, person.Id)
    	}
}
