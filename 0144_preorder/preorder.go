package preorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0, 4)
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
