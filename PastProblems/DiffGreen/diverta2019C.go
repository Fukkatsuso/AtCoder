// 答え見た

package main

import (
	"fmt"
)

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	BtoA := 0
	headB := 0
	tailA := 0
	containAB := 0
	for i := range s {
		head := s[i][0]
		tail := s[i][len(s[i])-1]
		if head == 'B' && tail == 'A' {
			BtoA++
		} else if head == 'B' {
			headB++
		} else if tail == 'A' {
			tailA++
		}
		// 1文字列にABが複数ある場合もある
		for j := 0; j < len(s[i])-1; j++ {
			if s[i][j:j+2] == "AB" {
				containAB++
			}
		}
	}

	ans := containAB
	if headB+tailA == 0 {
		ans += max(0, BtoA-1)
	} else {
		ans += BtoA + min(headB, tailA)
	}
	fmt.Println(ans)
}
