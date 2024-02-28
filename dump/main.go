package main

import "fmt"

func main() {
	len := 10
	leftView := make([]int, len)
	for ind, _ := range leftView {
		leftView[ind] = -1
	}
	fmt.Print(leftView)
}
