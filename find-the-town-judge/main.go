package main

func main()  {

}

func findJudge(N int, trust [][]int) int {
	rst := make([]int, N + 1)

	for _, t := range trust {
		rst[t[0]]--
		rst[t[1]]++
	}
	for index := 1; index < N + 1; index++ {
		if rst[index] == N - 1 {
			return index
		}
	}
	return -1
}