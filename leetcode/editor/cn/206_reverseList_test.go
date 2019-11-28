package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	// 输入: 1->2->3->4->5->NULL
	//输出: 5->4->3->2->1->NULL
	node := NewListNode(1)
	node.addNodeForValue(2).addNodeForValue(3).addNodeForValue(4).addNodeForValue(5)
	list := reverseList(node)
	assert.Equal(t, "5->4->3->2->1", list.String())
}
