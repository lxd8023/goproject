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

//106. 从中序与后序遍历序列构造二叉树【中等】
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	index := 0
	for inorder[index] != val {
		index++
	}
	node := &TreeNode{
		Val: val,
	}
	node.Left = buildTree1(inorder[:index], postorder[:index])
	node.Right = buildTree1(inorder[index+1:], postorder[index:len(postorder)-1])

	return node
}

//105. 从前序与中序遍历序列构造二叉树【中等】
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(inorder) <= 0 {
		return nil
	}
	val := preorder[0]
	index := 0
	for inorder[index] != val {
		index++
	}
	node := &TreeNode{
		Val: val,
	}

	node.Left = buildTree2(preorder[1:index+1], inorder[:index])
	node.Right = buildTree2(preorder[index+1:], inorder[index+1:])
	return node
}

//116. 填充每个节点的下一个右侧节点指针【中等】
//117. 填充每个节点的下一个右侧节点指针 II【中等】
func connect(root *Node) *Node {
	if root != nil {
		var nodes = []*Node{root.Left, root.Right}
		for len(nodes) > 0 {
			count := len(nodes)
			for i := 0; i < count; i++ {
				if nodes[i] != nil {
					if i < count-1 {
						nodes[i].Next = nodes[i+1]
					}
					if nodes[i].Left != nil {
						nodes = append(nodes, nodes[i].Left)
					}
					if nodes[i].Right != nil {
						nodes = append(nodes, nodes[i].Right)
					}
				}
			}
			nodes = nodes[count:]
		}
	}
	return root
}

//236. 二叉树的最近公共祖先【中等 】
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var resursive func(node, n, m *TreeNode) *TreeNode
	resursive = func(node, n, m *TreeNode) *TreeNode {
		if node != nil && node != n && node != q {
			left := resursive(node.Left, n, m)
			right := resursive(node.Right, n, m)

			if left == nil {
				return right
			}
			if right == nil {
				return left
			}
		}
		return node
	}

	return resursive(root, p, q)
}

//98. 验证二叉搜索树【中等】
func isValidBST(root *TreeNode) bool {
	var recursive func(node *TreeNode, min *TreeNode, max *TreeNode) bool
	recursive = func(node *TreeNode, min *TreeNode, max *TreeNode) bool {
		if node != nil {
			if min != nil && node.Val <= min.Val {
				return false
			}
			if max != nil && node.Val >= max.Val {
				return false
			}
			return recursive(node.Left, min, node) && recursive(node.Right, node, max)
		}
		return true
	}
	return recursive(root.Left, nil, root) && recursive(root.Right, root, nil)
}

//173. 二叉搜索树迭代器【中等】
type BSTIterator struct {
	Current int
	Length  int
	Nums    []int
}

func Constructor(root *TreeNode) BSTIterator {
	var res BSTIterator
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			node.Left = nil
			continue
		}
		res.Nums = append(res.Nums, node.Val)
		nodes = nodes[:len(nodes)-1]
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	res.Length = len(res.Nums)
	res.Current = -1
	return res
}

func (this *BSTIterator) Next() int {
	res := this.Nums[this.Current]
	this.Current++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.Current < this.Length
}
