## Go Leetcode Practice

My Go Leetcode Practice

## Leetcode editor For Goland

因为相同包下的方法只允许存在一个,遂采用不同包把题目分割开。

link:[solve golang package only have same func](https://github.com/shuzijun/leetcode-editor/issues/220)

CodeFileName: ${question.frontendQuestionId}/$!{question.frontendQuestionId}_${question.title}_test

```
package leetcode_cn_$!{question.frontendQuestionId}

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


${question.content}

${question.code}

func Test(t *testing.T) {
}

```