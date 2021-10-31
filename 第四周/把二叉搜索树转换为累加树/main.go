package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	return run(root)
}

func run(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = run(root.Right)
	root.Val += sum
	sum = root.Val
	root.Left = run(root.Left)
	return root
}

