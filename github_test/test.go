package main

import "fmt"

func main() {
	var a, b []int
	a = make([]int, 4)
	b = make([]int, 4)

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4

	b[0] = 1
	b[1] = 2
	b[2] = 2
	b[3] = 2
	fmt.Println(a, b)

	var c []int
	c = make([]int, 0)

	c = mergeSlices(a, b)

	fmt.Print(c)
}

func mergeSlices(s1 []int, s2 []int) []int {
	var slice []int
	slice = make([]int, 0)

	for i := 0; i < ; i++ {
		
	}
	
	return slice
}
