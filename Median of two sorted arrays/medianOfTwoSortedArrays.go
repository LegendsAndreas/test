// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	median := findMedianSortedArrays(nums1, nums2)

	fmt.Println("The median of the 2 sorted arrays are:", median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64
	numT := mergingArrays(nums1, nums2)

	numTLen := len(numT)
	var medianEle float64 // Which element of the array, the median is located in.

	if len(numT)%2 == 0 {
		var el1, el2 float64
		el1 = float64(numT[(numTLen/2)-1])
		el2 = float64(numT[(numTLen / 2)])
		median = (el1 + el2) / 2
	} else if len(numT)%2 != 0 {
		medianEle = float64(numTLen) / 2
		medianEle = math.Ceil(medianEle)
		median = float64(numT[int(medianEle)-1])
	} else {
		log.Fatal("This should not be possible wtf?")
	}

	return median
}

func mergingArrays(nums1 []int, nums2 []int) []int {
	var mergedArray []int

	return mergedArray
}
