package common

import "fmt"

//二叉树树形结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据Sequence构造TreeNode
func NewTreeNodeForSequence(nums []int) *TreeNode {
	return nil
}

func (root TreeNode) ConvertSequence() string {
	var result []int
	//var queue []int

	return fmt.Sprintf("%v", result)
}

func (root TreeNode) String() string {
	return root.ConvertSequence()
}
