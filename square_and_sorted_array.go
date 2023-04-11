package main

import "sort"

/*
Squares of a Sorted Array
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/

// -4, -2, 1, 3, 7, 9
// array given in sorted order
// make square root and sort it without sort function

func sortedSquares2(num []int) []int {

	var result []int
	for i := 0; i < len(num); i++ {
		sqrt := num[i] * num[i]
		result = append(result, sqrt)
	}
	sort.Ints(result) // This is the reason why it's O(nlogn)

	return result
}

func sortedSquares(nums []int) []int {
	var ptr1, ptr2 int = 0, len(nums) - 1
	//var currentIndex int = len(nums) - 1
	result := make([]int, len(nums))
	//for ptr2 >= ptr1 {
	for currentIndex := len(nums) - 1; currentIndex >= 0; currentIndex-- {
		val1, val2 := nums[ptr1], nums[ptr2]
		sqr1, sqr2 := val1*val1, val2*val2
		if sqr1 >= sqr2 {
			result[currentIndex] = sqr1
			ptr1++
		} else /* if sqr1 < sqr2 */ {
			result[currentIndex] = sqr2
			ptr2--
		}
		//currentIndex--
	}
	return result
}
