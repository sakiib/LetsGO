package main

import (
	"fmt"
)

const (
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE] *trieNode
	isWordEnd bool
}

var root *trieNode = &trieNode{}

func addWord(word string) {
    temp := root
    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'
        if temp.childrens[index] == nil {
            temp.childrens[index] = &trieNode{}
        }
        temp = temp.childrens[index]
    }
    temp.isWordEnd = true
} 

func findWord(word string) bool {
    temp := root
    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'
        if temp.childrens[index] == nil {
            return false
        }
        temp = temp.childrens[index]
    }
    return temp.isWordEnd
}

func main() {
	words := []string{"sam", "john", "tim", "jose", "rose", "cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		addWord(words[i])
	}
	
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose", "cat", "dog", "dogg", "roses", "rosess", "ans", "san"}

	for i := 0; i < len(wordsToFind); i++ {
		found := findWord(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
