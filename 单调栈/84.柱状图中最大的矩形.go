package theAlgorithm

import "fmt"

//  https://leetcode-cn.com/problems/largest-rectangle-in-histogram/submissions/

// 暴力，超时
func largestRectangleArea(heights []int) int {
	l := len(heights)

	var ans int

	for i := 0; i < l; i++ {
		w := 1
		h := heights[i]
		j := i - 1
		// 往左判断
		for j >= 0 && heights[j] >= h {
			w++
			j--
		}
		j = i + 1
		for j < l && heights[j] >= h {
			w++
			j++
		}
		ans = max(ans, w*h)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 单调栈

// 遍历每个柱体，
// 若当前的柱体高度大于等于栈顶柱体的高度，就直接将当前柱体入栈，
// 若当前的柱体高度小于栈顶柱体的高度，说明当前栈顶柱体找到了右边的第一个小于自身的柱体，
// 那么就可以将栈顶柱体出栈来计算以其为高的矩形的面积了。

func largestRectangleArea2(heights []int) int {
	// 增加哨兵
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	ans := 0
	var stack []int

	// 循环
	for i := 0; i < len(heights); i++ {
		if len(stack) > 0 && stack[len(stack)-1] > heights[i] {
			//计算栈顶元素
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			fmt.Println(h,i,stack[len(stack)-1])
			ans = max(ans, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return ans
}
