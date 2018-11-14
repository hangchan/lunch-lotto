package main

import "fmt"

type pair struct {
	name1     string
	name2     string
	day       int
	scheduled bool
}

func main() {
	people := []string{"John", "Jane", "Jim", "Steve"}
	pairs := makePairs(people)

	scheduleDay(pairs, people)
	for i, v := range pairs {
		fmt.Println(i, v)
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

	for i, v := range pairs {
		day = 1
		peopleMap = makePeopleMap(people)
		for ii, _ := range peopleMap {
			if v.scheduled == false && (ii == v.name1 || ii == v.name2) {
				pairs[i].day = day
				pairs[i].scheduled = true
				peopleMap[ii] = true
				peopleMap[ii] = true
				day++
			}
			fmt.Printf("index:%4v\tv.name1:%8v\tv.name2:%8v\tv.day:%4v\tv.scheduled:%8v\n", ii, pairs[i].name1, pairs[i].name2, pairs[i].day, pairs[i].scheduled)
		}
	}
	return pairs
}
