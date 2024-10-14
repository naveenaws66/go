package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for i, v := range(nums) {
        if i != len(nums) - 1 {
          for a, x := range(nums[i+1:]) {
              if v + x == target {
				  return []int{i, i+a+1}
				}
          }
        }
    }
	return []int{2, 2}
}

func main() {
	nums := []int{2, 11, 7, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

// {11, 2, 7, 15}
// {7, 15}

// 1 + 0 = 1 + 1 = 2
