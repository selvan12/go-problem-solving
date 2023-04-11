package main

/*
Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ { // Time: O(n)
		for j := i + 1; j < len(nums)-1; j++ { // Time: O(n)
			if nums[i]+nums[j] == target { // Space: O(1)
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int)     // Space: O(n)
	for index, value := range nums { // Time: O(n)
		if hValue, ok := hashMap[target-value]; ok {
			return []int{hValue, index}
		}
		hashMap[value] = index
	}
	return []int{}
}

// Time: O(n)
// Space: O(1)
// nums must be sorted!
func twoSum3(nums []int, target int) []int {
	ptr1, ptr2 := 0, len(nums)-1
	for ptr1 < ptr2 {
		sum := nums[ptr1] + nums[ptr2]
		if sum == target {
			return []int{ptr1, ptr2}
		} else if sum > target {
			ptr2--
		} else { // sum < target
			ptr2++
		}
	}

	return []int{}
}
