package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[string][]string)
	
	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	words := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(words)
	fmt.Println(result)
}
