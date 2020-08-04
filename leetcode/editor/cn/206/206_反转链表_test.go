package leetcode_cn_206

import (
	"github.com/stretchr/testify/assert"
	. "go-leetcode-practice/leetcode/editor/cn/common"
	"testing"
)

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1137 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedList(t *testing.T) {
	// 输入: 1->2->3->4->5->NULL
	//输出: 5->4->3->2->1->NULL
	node := NewListNode(1)
	node.AddNodeForValue(2).AddNodeForValue(3).AddNodeForValue(4).AddNodeForValue(5)
	list := reverseList(node)
	assert.Equal(t, "5->4->3->2->1", list.String())
}
