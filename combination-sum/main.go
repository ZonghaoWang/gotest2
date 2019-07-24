package main

import (
	"fmt"
)

func main()  {
	candidates := []int{2,3,6,7}
	target := 7
	rst := combinationSum(candidates, target)
	fmt.Printf("%v\n", rst)
}



//func combinationSum2(candidates []int, target int) [][]int {
//	sort.Slice(candidates, func(i int, j int) bool {
//		return candidates[i] > candidates[j]
//	})
//	fmt.Printf("%v\n", candidates)
//	return combinationSumSorted(candidates, target)
//}
func combinationSumSorted(candidates []int, target int) [][]int {
	rst := make([][]int, 0)
	if len(candidates) == 0 || candidates[len(candidates) - 1] > target {
		return nil
	}
	num := candidates[0]
	batchs := target / num
	for ind := batchs; ind >= 0; ind-- {
		fmt.Printf("ind = %d and num = %d and target = %d\n", ind, num, target - ind * num)

		pre := []int{}
		for i := 0; i < ind; i++ {
			pre = append(pre, num)
		}
		if target - ind * num == 0 && len(pre) > 0 {
			fmt.Printf("pre = %v\n", pre)
			rst = append(rst, pre)
		} else {
			tmp := combinationSumSorted(candidates[1:], target - ind * num)
			if len(tmp) > 0 {
				for _, t := range tmp {
					ll := append(t, pre...)
					fmt.Printf("ll = %v\n", ll)
					rst = append(rst, ll)
				}
			}
		}
	}
	return rst
}


func combinationSum(candidates []int, target int) (sols [][]int) {
	// sorting first
	sort(candidates)
	fmt.Println(candidates)
	return recur(candidates, target, 0, sols, []int{})
}

func recur(candidates []int, target int, cur int, sols [][]int, sol []int) [][]int {
	for i := cur; i < len(candidates); i++ {
		// 如果sum小於target，就繼續加下去
		// 大於就pop，因為已經超過，不需要再往下看更大的
		// 等於時的處理和大於相同，只是多了output
		switch sumAndJudge(sol, candidates[i], target) {
		case -1: // sum < target
			// 為了避免找出重複的解，recursive要帶入目前加到的項
			// 下一層的iterate就從該項開始(表示前面已經用過的項，往下不會再用到)
			sols = recur(candidates, target, i, sols, append(sol, candidates[i]))

		case 1:  // sum > target
			return sols

		case 0:  // sum == target
			sol = append(sol, candidates[i])
			return append(sols, append([]int{}, sol...))
		}
	}

	return sols
}

func sumAndJudge(sol []int, num int, target int) int {
	for i := 0; i < len(sol); i++ {
		target -= sol[i]
	}
	target -= num

	if target > 0 {
		return -1
	} else if target < 0 {
		return 1
	}
	return 0
}

func pop(slice []int, n int) []int {
	return slice[0:len(slice)-n]
}

func sort(candidates []int) {
	for i := 0; i < len(candidates)-1; i++ {
		for j := i+1; j > 0; j-- {
			if candidates[j] < candidates[j-1] {
				// swap
				candidates[j] = candidates[j] + candidates[j-1]
				candidates[j-1] = candidates[j] - candidates[j-1]
				candidates[j] = candidates[j] - candidates[j-1]
			}
		}
	}
}
