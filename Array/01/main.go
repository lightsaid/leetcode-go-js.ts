package main

/*
题目 #
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]

题目大意 #
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

*/
import (
	"fmt"
)
func main(){
	nums := []int{2, 7, 5, 11, 15, 23, 9, 11}
	target := 16
	fmt.Println("test1: ",test1(nums, target))
}

// 思路1：直接查找2个数的和=target
func test1(nums []int, target int) interface{}{
	l := len(nums)
	// 可能存在多个结果值
	res := map[string][]int{}
	for i:=0; i<l; i++{
		for j:=i+1; j<l; j++{
			if(nums[i] + nums[j] == target){
				var arr []int
				var s string
				arr = append(arr, i, j)
				s =  fmt.Sprintf("%d-%d", i, j)
				res[s] = arr
			}
		}
	}
	return res
}

