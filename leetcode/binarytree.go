package binarytree

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	Node struct {
		Val   int
		Left  *Node
		Right *Node
		Next  *Node
	}
)

//144. 二叉树的前序遍历【简单】
//方法：递归
func preorderRecursive(root *TreeNode) []int {
	var res []int
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		res = append(res, root.Val)
		preorderTraversal(root.Left)
		preorderTraversal(root.Right)
	}
	if root != nil {
		recursive(root)
	}
	return res
}

//方法：遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		var nodes = []*TreeNode{root.Right, root.Left}
		res = append(res, root.Val)
		for len(nodes) > 0 {
			node := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			if node != nil {
				res = append(res, node.Val)
				nodes = append(nodes, node.Right)
				nodes = append(nodes, node.Left)
			}
		}
	}
	return res
}

//94. 二叉树的中序遍历【简单】
//方法：递归
func inorderRecursive(root *TreeNode) []int {
	var res []int
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(root.Left)
			res = append(res, node.Val)
			recursive(root.Right)
		}
	}
	if root != nil {
		recursive(root)
	}
	return res
}

//方法：遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		var nodes = []*TreeNode{root}
		for len(nodes) > 0 {
			node := nodes[len(nodes)-1]
			if node != nil {
				if node.Left != nil {
					nodes = append(nodes, node.Left)
					node.Left = nil
					continue
				}
				res = append(res, node.Val)
				nodes = nodes[:len(nodes)-1]
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}
		}
	}
	return res
}

//145. 二叉树的后序遍历【简单】
//方法：递归
func postorderRecursive(root *TreeNode) []int {
	var res []int
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(node.Left)
			recursive(node.Right)
			res = append(res, node.Val)
		}
	}
	recursive(root)
	return res
}

//方法：遍历
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		var nodes = []*TreeNode{root}
		for len(nodes) > 0 {
			node := nodes[len(nodes)-1]
			if node != nil {
				if node.Right != nil {
					nodes = append(nodes, node.Right)
					node.Right = nil
					continue
				}
				if node.Left != nil {
					nodes = append(nodes, node.Left)
					node.Left = nil
					continue
				}
				res = append(res, node.Val)
				nodes = nodes[:len(nodes)-1]
			}
		}
	}
	return res
}

//104. 二叉树的最大深度【简单】
func maxDepth(root *TreeNode) int {
	var res int
	var recursive func(node *TreeNode, n int)
	recursive = func(node *TreeNode, n int) {
		if node != nil {
			if n > res {
				res = n
			}
			recursive(node.Left, n+1)
			recursive(node.Right, n+1)
		}
	}
	recursive(root, res+1)
	return res
}

//101. 对称二叉树【简单】
func isSymmetric(root *TreeNode) bool {
	var recursive func(left *TreeNode, right *TreeNode) bool
	recursive = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if right == nil || left == nil {
			return false
		}
		return left.Val == right.Val && recursive(left.Left, right.Right) && recursive(left.Right, right.Left)
	}
	return recursive(root.Left, root.Right)
}

//112. 路径总和【简单】
func hasPathSum(root *TreeNode, targetSum int) bool {
	var resursive func(node *TreeNode, n int) bool
	resursive = func(node *TreeNode, n int) bool {
		if node == nil {
			return false
		}
		if node.Left == nil && node.Right == nil && node.Val == n {
			return true
		}
		return resursive(node.Left, n-node.Val) || resursive(node.Right, n-node.Val)
	}
	return resursive(root, targetSum)
}
