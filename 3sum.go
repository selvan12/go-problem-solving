package main

import "sort"

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/
func threeSum(nums []int) [][]int {
	n := len(nums)

	// hashMap used here to avoid duplicate
	hashMap := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {

				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					hashMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	var result [][]int
	for _, triplet := range hashMap {
		result = append(result, triplet)
	}

	return result
}

func threeSum2(items []int) [][]int {
	l := [][]int{}
	n := len(items)
	sort.Ints(items)

	// Remove the duplicates from the items array.
	for i := range items {
		if i > 0 && items[i-1] == items[i] {
			continue
		}

		// Now take the two pointer approach to solve the linear array problem.
		j, k := i+1, n-1
		for j < k {
			sum := items[i] + items[j] + items[k]
			// if sum  is greater than target means the larger
			// value(from right as items is sorted i.e, k at right)
			// is taken and it is not able to sum up to the target
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				l = append(l, []int{items[i], items[j], items[k]})
				j++
				for items[j-1] == items[j] && j < k {
					j++
				}
			}

		}
	}
	return l
}
