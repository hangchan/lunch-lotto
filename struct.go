package main

import (
	"fmt"
	"sort"
)

type pair struct {
	day       int
	name1     string
	name2     string
	scheduled bool
}

func main() {
	people := []string{"Jim", "John", "James", "Bill", "Ted"}
	pairs := makePairs(people)

	scheduleDay(pairs, people)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].day < pairs[j].day
	})
	for _, v := range pairs {
		fmt.Printf("Day: %v\t Pair: %v-%v\t\n", v.day, v.name1, v.name2)
	}

}

func newPair(first string, second string) pair {
	var p pair

	if first != second {
		p = pair{name1: first, name2: second}
		return p
	}

	return p
}

func makePairs(people []string) []pair {
	var pairs []pair

	for i, _ := range people {
		for j := i + 1; j < len(people); j++ {
			p := newPair(people[i], people[j])
			pairs = append(pairs, p)
		}
	}

	return pairs
}

func makePeopleMap(people []string) map[string]bool {
	// make peopleMap
	peopleMap := make(map[string]bool)

	for _, v := range people {
		peopleMap[v] = false
	}

	return peopleMap
}

func scheduleDay(pairs []pair, people []string) []pair {
	day := 1
	peopleMap := makePeopleMap(people)
	pairsScheduled := 0

	for i := 0; i < 10; i++ {
		peopleMap = makePeopleMap(people)
		for i, _ := range peopleMap {
			for ii, vv := range pairs {
				if vv.scheduled == false && (i == vv.name1 || i == vv.name2) && (peopleMap[vv.name1] == false && peopleMap[vv.name2] == false) {
					pairs[ii].day = day
					pairs[ii].scheduled = true
					peopleMap[vv.name1] = true
					peopleMap[vv.name2] = true
					pairsScheduled = pairsScheduled + 1
				}
			}
		}
		day++
		if pairsScheduled == len(pairs) {
			break
		}
	}

	return pairs
}
