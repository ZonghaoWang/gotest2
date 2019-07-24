package main

func main()  {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ls, lp := doLongestUnivaluePath(root.Left, root.Val)
	rs, rp := doLongestUnivaluePath(root.Right, root.Val)
	return intMax(ls + rs + 1, intMax(lp, rp)) - 1
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func doLongestUnivaluePath(root *TreeNode, target int) (s, p int) {
	if root == nil {
		return 0, 0
	}
	var ms, mp int

	ls, lp := doLongestUnivaluePath(root.Left, root.Val)
	rs, rp := doLongestUnivaluePath(root.Right, root.Val)
	if root.Val != target {
		ms = 0
	} else {
		ms = intMax(ls, rs) + 1
	}
	mp = intMax(intMax(lp, rp), ls + rs + 1)
	return ms, mp
}