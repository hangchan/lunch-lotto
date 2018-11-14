package main

import (
	"fmt"
)

type pair struct{
	name1 string
	name2 string
}

func main() {
	people := []string{"John", "Jane", "Jim", "Steve", "Bill"}
	// get a list of pairs
	// pairs := makePairs(people)
	// schedule day for pairs
	//scheduleDay(pairs)
	printPairs(people)
}

func newPair(first string, second string) pair {
		var p pair
		if first != second {
			p = pair{name1: first, name2: second}
			return p
		}
		return p
}

func printPairs (people []string) {
	for i, _ := range people {
		for j := i+1; j < len(people); j++ {
			p := newPair(people[i], people[j])
			fmt.Printf("Pair:\t%v\n", p)
		}
	}
}
