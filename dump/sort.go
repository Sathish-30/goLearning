package main

import (
	"fmt"
	"sort"
)

// Where to sort a struct we need to implement the sort.Interface methods and use sort.Sort method
type Arr []int

func (arr Arr) Len() int {
	return len(arr)
}

func (arr Arr) Less(i, j int) bool {
	return arr[i] > arr[j]
}

func (arr Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sortingImplementation() {
	arr := []int{11, 2, 32, 4, 5, 64, 7, 8}
	sort.Sort(Arr(arr))
	fmt.Println(arr)
}
