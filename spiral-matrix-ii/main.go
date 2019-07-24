package main

import "fmt"

func main()  {
	fmt.Println(generateMatrix(3))
}

type Direction int
const (
	_ Direction = iota
	START
	RIGHT
	DOWN
	LEFT
	TOP
)


func generateMatrix(n int) [][]int {
	rst := make([][]int, n)
	for ind := 0; ind < n; ind++ {
		rst[ind] = make([]int, n)
	}
	l, t := 0, 0
	r, b := n - 1, n - 1
	x, y := 0, 0
	dire := START
	status := false
	cnt := 1
	for l <= r && t <= b {
		switch dire {
		case START:
			rst[x][y] = cnt
			if r == y {
				if x == b {
					return rst
				}
				x++
				dire = DOWN
			} else {
				y++
				dire = RIGHT
			}
			status = true
		case RIGHT:
			if status {
				if y == l {
					y++
					l++
					status = false
					continue
				}
			}
			cnt++
			rst[x][y] = cnt
			if y == r {
				dire = DOWN
				status = true
				continue
			}
			y++
		case DOWN:
			if status {
				if x == t {
					x++
					t++
					status = false
					continue
				}
			}
			cnt++
			rst[x][y] = cnt
			if x == b {
				dire = LEFT
				status = true
				continue
			}
			x++
		case LEFT:
			if status {
				if y == r {
					y--
					r--
					status = false
					continue
				}
			}
			cnt++
			rst[x][y] = cnt
			if y == l {
				dire = TOP
				status = true
				continue
			}
			y--
		case TOP:
			if status {
				if x == b {
					x--
					b--
					status = false
					continue
				}
			}
			cnt++
			rst[x][y] = cnt
			if x == t {
				status = true
				dire = RIGHT
				continue
			}
			x--
		}
	}
	return rst
}

