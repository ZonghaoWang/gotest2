package main
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func main()  {

}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right != nil {
		return hasPathSum(root.Right, sum - root.Val)
	}
	if root.Right == nil && root.Left != nil {
		return hasPathSum(root.Left, sum - root.Val)
	}
	if root.Right != nil && root.Left != nil {
		return hasPathSum(root.Right, sum - root.Val) || hasPathSum(root.Left, sum - root.Val)
	}
	return sum == root.Val
}