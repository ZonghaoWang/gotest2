package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main()  {

}

func pathSum(root *TreeNode, sum int) [][]int {
	rst := make([][]int, 0)
	current := make([]int, 0)
	doPathSum(root, sum, &rst, &current)
	return rst
}

func doPathSum(root *TreeNode, remain int, rst *[][]int, current *[]int)  {
	if root == nil {
		return
	}
	if remain == root.Val && root.Left == nil && root.Right == nil {
		*current = append(*current, remain)
		cc := make([]int, len(*current))
		copy(cc, *current)
		*rst = append(*rst, cc)
		*current = (*current)[:len(*current) - 1]
		return
	}
	*current = append(*current, root.Val)
	if root.Right != nil {
		doPathSum(root.Right, remain - root.Val, rst, current)
	}
	if root.Left != nil {
		doPathSum(root.Left, remain - root.Val, rst, current)
	}
	*current = (*current)[:len(*current) - 1]
}