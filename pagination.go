package main

import "fmt"

func fetchItemsToDisplay(items [][]string, sortParameter, sortOrder, itemsPerPage, pageNumber int32) string {

	return "Hello World"
}

func main() {

	items := [][]string{
		{"item 1", "10", "15"},
		{"item 1", "3", "4"},
		{"item 1", "17", "18"},
	}

	var sortParameter, sortOrder, itemsPerPage, pageNumber int32 = 1, 0, 2, 1

	fmt.Println(fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber))
}
