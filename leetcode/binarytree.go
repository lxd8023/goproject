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
