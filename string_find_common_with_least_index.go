package main

/*
Minimum Index Sum of Two Lists:
Given two arrays of strings list1 and list2, find the common strings with the least index sum.
A common string is a string that appeared in both list1 and list2.
A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j should be the minimum value among all the other common strings.
Return all the common strings with the least index sum. Return the answer in any order.

Example 1:
Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only common string is "Shogun".

Example 2:
Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.

Example 3:
Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
Output: ["sad","happy"]
Explanation: There are three common strings:
"happy" with index sum = (0 + 1) = 1.
"sad" with index sum = (1 + 0) = 1.
"good" with index sum = (2 + 2) = 4.
The strings with the least index sum are "sad" and "happy".
*/

func findRestaurant(list1 []string, list2 []string) []string {
	count := 1000 //math.MaxInt
	result := make([]string, 1)

	for i, str1 := range list1 {
		for j, str2 := range list2 {
			if str1 == str2 {
				if i+j < count {
					count = i + j
					result[0] = str1
				} else if i+j == count {
					result = append(result, str1)
				}
			}
		}
	}

	return result
}

func findRestaurant2(list1 []string, list2 []string) []string {
	a := map[string]int{}
	for i, val := range list1 {
		for j, vall := range list2 {
			if val == vall {
				a[vall] = j + i
			}
		}
	}

	minSum := len(list1) + len(list2) - 2
	for _, sum := range a {
		if sum < minSum {
			minSum = sum
		}
	}

	res := []string{}
	for i, sum := range a {
		if sum == minSum {
			res = append(res, i)
		}
	}
	return res
}
