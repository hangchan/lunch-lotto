package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	people := []string{"John", "Jane", "Jim", "Steve", "Bill"}
	// get a list of pairs
	pairs := makePairs(people)
	for i,_ := range pairs {
		fmt.Printf("Pair:\t%v\n", i)
	}
	// schedule day for pairs
}

func makePairs(people []string) map[string]int {
	pairsMap := make(map[string]int)
	pair := []string{}
	// make pairs from people
	for i,_ := range people {
		for j,_ := range people{
			if people[i] != people[j] {
				// fmt.Printf("OuterLoop:\t%v\tInnerLoop:\t%v\n", people[i], people[j])
				pair = []string{people[i], people[j]}
				sort.Strings(pair)
				pairStr := strings.Join(pair, "-")
				// fmt.Printf("Type: %T\n", pair)
				// fmt.Println("PairStr:", pairStr)
				pairsMap[pairStr] = 0
			}
		}
	}
	// pairs = sort.StringSlice(pairs)
	return pairsMap
}

func sliceMatch (sliceStr []string, str string) bool {
	for _, v := range sliceStr {
		if str == v {
			fmt.Println("SliceMatch:", sliceStr, str)
			return true
		}
	}
	return false
}