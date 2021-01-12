package leetcode_cn_870

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

//给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
//
// 返回 A 的任意排列，使其相对于 B 的优势最大化。
//
//
//
// 示例 1：
//
// 输入：A = [2,7,11,15], B = [1,10,4,11]
//输出：[2,11,7,15]
//
//
// 示例 2：
//
// 输入：A = [12,24,8,32], B = [13,25,32,11]
//输出：[24,32,8,12]
//
//
//
//
// 提示：
//
//
// 1 <= A.length = B.length <= 10000
// 0 <= A[i] <= 10^9
// 0 <= B[i] <= 10^9
//
// Related Topics 贪心算法 数组
// 👍 103 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func advantageCount(A []int, B []int) []int {
	aziped := zip(A)
	bziped := zip(B)
	ans := make([]int, len(A))
	for p, r, l := len(A)-1, len(A)-1, 0; p >= 0; p-- {
		if aziped[r][1] > bziped[p][1] {
			ans[bziped[p][0]] = aziped[r][1]
			r--
		} else {
			ans[bziped[p][0]] = aziped[l][1]
			l++
		}
	}
	return ans
}

func zip(nums []int) [][]int {
	N := len(nums)
	ans := make([][]int, N)
	for i := range nums {
		ans[i] = make([]int, 2)
		ans[i][0] = i
		ans[i][1] = nums[i]
	}

	sort.SliceStable(ans, func(i, j int) bool {
		if ans[i][1] == ans[j][1] {
			return ans[i][0] < ans[j][0]
		}
		return ans[i][1] < ans[j][1]
	})
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func Test(t *testing.T) {
	assert.Equal(t, []int{2, 11, 7, 15}, advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	assert.Equal(t, []int{24, 32, 8, 12}, advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
