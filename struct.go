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
	people := []string{"John", "Jane", "Jim", "Bill", "Joe"}
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
			//fmt.Printf("name1:%8v\tname2:%8v\n", p.name1, p.name2)
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
				//fmt.Println("Day", day)
				if vv.scheduled == false && (i == vv.name1 || i == vv.name2) && (peopleMap[vv.name1] == false && peopleMap[vv.name2] == false) {
					pairs[ii].day = day
					pairs[ii].scheduled = true
					peopleMap[vv.name1] = true
					peopleMap[vv.name2] = true
					pairsScheduled = pairsScheduled + 1
				}
				//fmt.Printf("index:%5v\tv.name1:%8v\tv.name2:%8v\tv.day:%4v\tv.scheduled:%8v\n", i, pairs[ii].name1, pairs[ii].name2, pairs[ii].day, pairs[ii].scheduled)
				//fmt.Println(peopleMap)
			}
		}
		day++
		if pairsScheduled == len(pairs) {
			break
		}
	}

	return pairs
}
