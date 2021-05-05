/*
From: https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/a/divide-and-conquer-algorithms

Divide-and-conquer, breaks a problem into subproblems that are similar to the original problem, recursively solves the subproblems, and finally combines the solutions to the subproblems to solve the original problem. Because divide-and-conquer solves subproblems recursively, each subproblem must be smaller than the original problem, and there must be a base case for subproblems

Solve the following problem with a divide and conquer approach.

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

PEDAC

Problem:
Create contiguous subarrays, starting at each element in the input array. Sum the values in the subarray and return the largest sum.

Based on the examples provided, the initial input array should also be included.

Data structures:
Will reslice the input slice

Algorithm:
- Sum the initial slice
- create a base case of an empty input slice
	- return `sum`
- create a sum helper method
- first iteration will be subarray from indexes 0 to slice length - 1, then slice length - 2 etc until subarray is empty
- also need to iterate in the other direction, start at index 1 to the end, then index 2 to the end etc

*/

package main

import "fmt"

func maxSubArray(nums []int) []int {
	subs := []int{}
	if len(nums) == 0 {
		return subs
	}

	sum := sumOfSubarray(nums)

	subs = append(subs, sum)
	subs = append(subs, maxSubArray(nums[1:])...)
	subs = append(subs, maxSubArray(nums[:(len(nums)-1)])...)

	return subs
}

func sumOfSubarray(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
