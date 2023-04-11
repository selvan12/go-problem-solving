package main

import "math"

/*
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.

Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var findDepth func(root *TreeNode) int
	// define inline function
	findDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := findDepth(node.Left)
		right := findDepth(node.Right)
		ans = int(math.Max(float64(ans), float64(left+right+1)))

		return int(math.Max(float64(left), float64(right))) + 1

	}

	findDepth(root)

	return ans - 1
}
