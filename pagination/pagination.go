package main

import (
	"fmt"
	"sort"
)

func pagination(itemsPerPage, pageNumber int32, items []string) {
	totalPages := len(items) / int(itemsPerPage)
	restDivision := len(items) % int(itemsPerPage)
	result := make([][]string, totalPages)

	if restDivision == 0 {
		for i := 0; i < totalPages; i++ {
			result[i] = append(result[i], items[i])
		}
	}
}

func sortItems(sortOrder int32, items []string) []string {

	if sortOrder == 0 {
		sort.Sort(sort.StringSlice(items))
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(items)))
	}

	return items
}

func fetchItemsToDisplay(items [][]string, sortParameter, sortOrder, itemsPerPage, pageNumber int32) [][]string {

	var n, nNames int = len(items), 0
	var arrItems []string

	for i := 0; i < n; i++ {
		for _, s := range items {
			arrItems = append(arrItems, s[i])
			nNames++
			if nNames == n {
				break
			}
		}
		if nNames == n {
			break
		}
	}

	arrItems = sortItems(sortOrder, arrItems)

	pagination(itemsPerPage, pageNumber, arrItems)

	return items
}

func main() {

	// [name, relevance, price]
	items := [][]string{
		{"jose", "10", "15"},
		{"adalto", "3", "4"},
		{"zoraide", "17", "18"},
		{"dirce", "10", "15"},
		{"ana", "3", "4"},
		{"rosilda", "17", "18"},
	}

	var sortParameter, sortOrder, itemsPerPage, pageNumber int32 = 1, 0, 2, 1

	fmt.Println(fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber))
}
