package main

func main()  {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	var total int
	doPathSum(root, sum, &total)
	return total
}

func doPathSumBeginWith(root *TreeNode, remain int, total *int) {
	if root == nil {
		return
	}
	if root.Val == remain {
		*total++
	}
	doPathSumBeginWith(root.Left, remain - root.Val, total)
	doPathSumBeginWith(root.Right, remain - root.Val, total)
}

func doPathSum(root *TreeNode, remain int, total *int)  {
	if root == nil {
		return
	}

	doPathSum(root.Left, remain - root.Val, total)
	doPathSum(root.Right, remain - root.Val, total)
	doPathSumBeginWith(root, remain, total)
}