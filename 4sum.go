package main

import (
	"encoding/json"
	"sort"
)

/*
4Sum
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	uniqueMap := make(map[string]bool, 0)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					row := []int{nums[i], nums[j], nums[l], nums[r]}
					b, _ := json.Marshal(row)
					if _, ok := uniqueMap[string(b)]; !ok {
						uniqueMap[string(b)] = true
						res = append(res, row)
					}
					l++
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res

}
