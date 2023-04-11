package main

/*
Pow(x, n)
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:
Input: x = 2.00000, n = 10
Output: 1024.00000

Example 2:
Input: x = 2.10000, n = 3
Output: 9.26100

Example 3:
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	result := 1.0
	for n > 0 {
		result *= x
		n--
	}
	return result
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow2(x, -n)
	}
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}
	if n%2 == 0 {
		temp := myPow2(x, n/2)
		return temp * temp
	}
	return x * myPow2(x, n-1)
}
