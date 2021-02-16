package main

import (
	"fmt"
	"sort"
)

func fetchItemsToDisplay(items [][]string, sortParameter, sortOrder, itemsPerPage, pageNumber int32) [][]string {

	var n, nNames int = len(items), 0
	//var finalItensFetched [][]string
	var arrName /*, arrRelevance, arrPrice*/ []string

	for i := 0; i < n; i++ {
		for _, s := range items {
			arrName = append(arrName, s[i])
			nNames++
			if nNames == n {
				break
			}
		}
		if nNames == n {
			break
		}
	}

	if sortOrder == 0 {
		sort.Sort(sort.StringSlice(arrName))
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(arrName)))
	}

	fmt.Println(n)
	fmt.Println(arrName)

	return items
}

func main() {

	// [name, relevance, price]
	items := [][]string{
		{"jose", "10", "15"},
		{"adalto", "3", "4"},
		{"zoraide", "17", "18"},
	}

	var sortParameter, sortOrder, itemsPerPage, pageNumber int32 = 1, 0, 2, 1

	fmt.Println(fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber))
}
