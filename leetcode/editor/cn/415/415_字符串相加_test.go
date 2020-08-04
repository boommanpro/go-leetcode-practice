package leetcode_cn_415

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 注意：
//
//
// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
// Related Topics 字符串
// 👍 207 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func addStrings(num1 string, num2 string) string {
	carry := 0
	ans := ""
	for l1, l2 := len(num1)-1, len(num2)-1; l1 >= 0 || l2 >= 0 || carry != 0; l1, l2 = l1-1, l2-1 {
		var v1, v2 int
		if l1 >= 0 {
			v1 = int(num1[l1] - '0')
		}
		if l2 >= 0 {
			v2 = int(num2[l2] - '0')
		}
		sum := v1 + v2 + carry
		carry = sum / carry
		ans = strconv.Itoa(sum%10) + ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func Test(t *testing.T) {
	assert.Equal(t, "123456", addStrings("123454", "2"))
	assert.Equal(t, "10", addStrings("9", "1"))
}
