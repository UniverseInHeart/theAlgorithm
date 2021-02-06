package main

import (
	"fmt"
	"sort"
)
func main(){
	fmt.Println("hello")
	
}

func maximumProduct(nums []int) int {

	if len(nums) == 3 {
		a := nums[0] * nums[1] * nums[2]
		return a
	}

	sort.Ints(nums)

	if nums[len(nums)-1] < 0 {
		return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	}
	if nums[len(nums)-1] == 0 {
		return 0
	}

	if nums[0] < 0 && nums[1] < 0 {
		if nums[0]*nums[1] > nums[len(nums)-2]*nums[len(nums)-3] {
			return nums[0] * nums[1] * nums[len(nums)-1]
		}
	}
	return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]

}

func maximumTime(time string) string {
	src := []rune(time)

	a := byte(src[0])
	b := byte(src[1])
	if a == 63 {
		if b > 54 {
			src[0] = 49
		} else {
			src[0] = 50
		}
	}

	if b == 63 {
		if a < 50 {
			src[1] = 57
		} else {
			src[1] = 51
		}
	}

	c := byte(src[3])
	if c == 63 {
		src[3] = 53
	}

	d := byte(src[4])
	if d == 63 {
		src[4] = 57
	}
	fmt.Print(src)

	return string(src)
}
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	size := 1
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > last {
			size++
		} else {
			size = 1
		}
		last = nums[i]
		if size > ans {
			ans = size
		}
	}
	return ans
}

func pivotIndex(nums []int) int {

	ans := -1
	for j := 0; j < len(nums); j++ {
		lll := 0
		for i := 0; i < j; i++ {
			lll += nums[i]
		}
		for i := j + 1; i < len(nums); i++ {
			lll -= nums[i]
		}
		if lll == 0 {
			ans = j
			break
		}
	}
	return ans
}

func pivotIndex2(nums []int) int {

	list := make([]int, len(nums))
	for j := 0; j < len(nums); j++ {
		list[j] = nums[j]
	}

	for j := 1; j < len(nums); j++ {
		list[j] = list[j] + list[j-1]
	}
	if list[len(nums)-1] == list[0] {
		return 0
	}

	for j := 1; j < len(nums); j++ {
		if list[j-1] == list[len(nums)-1]-list[j] {
			return j
		}
	}
	if list[len(nums)-2] == 0 {
		return 0
	}
	return -1
}

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	n := 0
	for i := lowLimit; i < highLimit+1; i++ {
		temp := i
		for temp > 0 {
			n += temp % 10
			temp /= 10
		}
		m[n]++
		n = 0
	}
	ans := 0
	for _, v := range m {
		ans = max(ans, v)
	}

	return ans

}

// 从左右断点出发去组装，因为元素不同所以能使用这个方法
func restoreArray(adjacentPairs [][]int) []int {
	ans := make([]int, len(adjacentPairs)+1)
	// 找出最右边和最左边，因为其只能出现一次
	m := make(map[int]int)
	m2 := make(map[int][]int)
	for _, v := range adjacentPairs {
		m[v[0]]++
		m[v[1]]++
		if m2[v[0]] == nil {
			m2[v[0]] = []int{v[1]}
		} else {
			m2[v[0]] = append(m2[v[0]], v[1])
		}

		if m2[v[1]] == nil {
			m2[v[1]] = []int{v[0]}
		} else {
			m2[v[1]] = append(m2[v[1]], v[0])
		}

	}
	var left int
	for k, v := range m {
		if v == 1 {
			left = k
			break
		}
	}
	ans[0] = left
	for i := 1; i < len(adjacentPairs)+1; i++ {

		temp := m2[left]
		for _, v := range temp {
			if v != ans[i-2] || i == 1 {
				left = v
			}
		}

		ans[i] = left
	}
	return ans
}

func checkPartitioning(s string) bool {
	runes := []rune(s)

	for i := 1; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			one := runes[:i+1]
			two := runes[i : j+1]
			three := runes[j:]
			fmt.Println(one, two, three)
			if isReverse(one) && isReverse(two) && isReverse(three) {
				return true
			}
		}
	}
	return false

}
func isReverse(runes []rune) bool {
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		if runes[from] != runes[to] {
			return false
		}
	}
	return true
}

func canEat(candiesCount []int, queries [][]int) []bool {
	// 前缀和
	sum := make([]int, len(candiesCount))

	sum[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++ {
		sum[i] = sum[i-1] + candiesCount[i]
	}
	fmt.Println(sum)
	ans := make([]bool, 0)

	// 0  1  2  3  4
	//
	for _, v := range queries {
		favoriteType := v[0]
		b := sum[favoriteType]
		s := 0
		if favoriteType > 0 {
			s = sum[favoriteType-1]
		}
		ss := v[1] + 1
		bb := v[1]*v[2] + v[2]

		if b < ss || s >= bb {
			ans = append(ans, false)
		} else {
			ans = append(ans, true)
		}
	}

	return ans
}

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1 > k+maxCnt {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
