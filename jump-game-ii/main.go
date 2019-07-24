package main

import "fmt"

func main()  {
	a := []int{2,3,1,1,4}
	fmt.Println(jump(a))
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var times, left, right int
	max := 0
	for {
		times++
		for i := left; i <= right; i++ {
			max = intMax(max, nums[i] + i)
		}
		if max >= len(nums) - 1 {
			break
		}
		left = right + 1
		right = max
	}
	return times
}
