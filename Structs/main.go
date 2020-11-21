package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Id   int
}

// in-case we have another struct reference 
type Person struct {
	Name string
	Id   int
	Language *Language
}

type Language struct {
    Name string
    Id int
    Exp string
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
	fmt.Println(s)
	
	// Sort by Id, keeping original order or equal elements.
    	sort.SliceStable(personList, func(i, j int) bool {
       		return personList[i].Id < personList[j].Id
    	})
	
    	for _, person := range personList {
		fmt.Println(person.Name, person.Id)
    	}
	
	// struct Person have another reference to struct Language
	person1 := Person{Name: "sakib", Id: 1, Language: &Language{Name: "C++", Id: 1, Exp: "Good"}}
	fmt.Println(person1.Name, person1.Id, person1.Language.Name, person1.Language.Id, person1.Language.Exp)
}
