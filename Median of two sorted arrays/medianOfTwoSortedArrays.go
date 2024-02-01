// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	// Our two arrays that we wish to merge
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 3, 4, 4, 4}

	median := findMedianSortedArrays(nums1, nums2)

	// Prints what the median is with the merged arrays.
	fmt.Println("The median of the 2 sorted arrays are:", median)
}

/*
*
Finds the median value from the two sorted arrays and returns it as a float64 value.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64
	mergedNums := mergingArrays(nums1, nums2)

	mergedNumsLen := len(mergedNums)
	var medianEle float64 // Which element of the array, the median is located in.

	// In case the length of mergedNums is even, odd or cursed.
	if mergedNumsLen%2 == 0 {
		var el1, el2 float64
		el1 = float64(mergedNums[(mergedNumsLen/2)-1])
		el2 = float64(mergedNums[(mergedNumsLen / 2)])
		median = (el1 + el2) / 2
	} else if mergedNumsLen%2 != 0 {
		medianEle = float64(mergedNumsLen) / 2
		medianEle = math.Ceil(medianEle)
		median = float64(mergedNums[int(medianEle)-1])
	} else {
		log.Fatal("This should not be possible wtf?")
	}

	return median
}

/*
*
Merging the two sorted arrays into one slice.
*/
func mergingArrays(nums1 []int, nums2 []int) []int {
	var mergedArray []int
	mergedArray = make([]int, 0)

	// Adds the elements from the first array into mergedArray
	for i := 0; i < len(nums1); i++ {
		mergedArray = append(mergedArray, nums1[i])
	}

	// Adds the elements from the second array into mergedArray
	for i := 0; i < len(nums2); i++ {
		mergedArray = append(mergedArray, nums2[i])
	}

	// Just a Golang sorting option. Will probably implement my own sorting function.
	sort.Ints(mergedArray)

	//fmt.Println(mergedArray)

	return mergedArray
}
