package leetcode_cn_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
// 👍 8800 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, v := range nums {
		search := target - v
		if p, ok := dict[search]; ok {
			return []int{p, i}
		}
		dict[v] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
}
