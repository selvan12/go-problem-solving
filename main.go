package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//var head *ListNode

func createList(arr []int) *ListNode {
	var head *ListNode
	for _, val := range arr {
		list := &ListNode{val, nil}
		if head == nil {
			head = list
		} else {
			temp := head
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = list
		}
	}
	return head
}

func main() {
	log.Println("Coding exercise")

	// arr := []int{1, 8, 3, 8, 1}
	// arr := []int{9, 1, 1}
	// if checkPalindrome((arr)) {
	// if checkPalindromeStack((arr)) {

	//head := createList([]int{1, 8, 3, 8, 1})
	//head := createList([]int{9, 1, 1})

	// if isPalindrome(head) {
	// 	fmt.Println("given array is palindrome")
	// } else {
	// 	fmt.Println("given array is not a palindrome")
	// }

	//fmt.Println("isPalindromeString: ", isPalindromeString("MADAM"))

	// list1 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	// list2 := createList([]int{9, 9, 9, 9})
	// addTwoNumbers(list1, list2)

	//printNumbers(5)
	//printNumbersRecursive(5)

	// printPyramid()
	// printPyramid2()
	// printPyramidRev()

	// var stack Stack
	// stack.push(3)
	// stack.push(4)
	// stack.push(5)
	// for len(stack) > 0 {
	// 	fmt.Println(stack.pop())
	// 	ret, len := stack.isEmpty()
	// 	fmt.Println("isEmpty: ", ret, " len:", len)
	// }

	//mystry(6)

	//queue()

	//httpServer()

	//arr := []int{0, 0, 0, 2, 0, 0}
	//zeroFilledSubarray(arr)

	// s := "abcabcbb"
	// s := "bbbbb"
	// s := "pwwkew"
	// fmt.Println("lengthOfLongestSubstring ", lengthOfLongestSubstring(s))

	//fmt.Println("reverse_integer: ", reverse_integer(123))

	// mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println("findDiagonalOrder: ", findDiagonalOrder(mat))

	// fmt.Println("reverseWordInString: ", reverseWordInString("hello this is the test string"))

	// headA := createList([]int{4, 1})
	// headB := createList([]int{5, 6, 1})
	// intersec := createList([]int{8, 4, 5})
	// headA.Next.Next = intersec
	// headB.Next.Next.Next = intersec
	// out := getIntersectionNode2(headA, headB)
	// fmt.Println("Otput: ", out)

	// str1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	// str2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	// str1 := []string{"happy", "sad", "good"}
	// str2 := []string{"sad", "happy", "good"}
	// fmt.Println("findRestaurant: ", findRestaurant2(str1, str2))

	// s := "MCMXCIV" //"LVIII" //"III"
	// fmt.Println("romanToInt: ", romanToInt(s))

	// fmt.Println("Out: ", sortedSquares([]int{-4, -2, 1, 3, 7, 9}))

	//convertTo24HourFormat2("11:00:00 PM")

	//fmt.Println("isPalindromeNumber: ", isPalindromeNumber2(-121))

	// fmt.Println("Equivalent Binary Trees: ", Same(tree.New(1), tree.New(1)))
	// fmt.Println("Equivalent Binary Trees: ", Same(tree.New(1), tree.New(2)))
	// fmt.Println("Equivalent Binary Trees: ", Same(tree.New(10), tree.New(10)))

	//camelCaseWordCount()

	// s := "PAYPALISHIRING"
	// fmt.Println("result: ", convert(s, 3))

	// fmt.Println("myPow result: ", myPow2(2.00000, -2))

	// nums := []int{2, 7, 11, 15}
	// target := 9
	// fmt.Println("twoSum: ", twoSum3(nums, target))

	// arr := []int{-1, 0, 1, 2, -1, -4}
	// arr := []int{0, 1, 1}
	// fmt.Println("threeSum: ", threeSum(arr))

	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("fourSum: ", fourSum(nums, target))
}
