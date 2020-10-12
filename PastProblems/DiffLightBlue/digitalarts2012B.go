package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func num(c byte) int {
	return int(c-'a') + 1
}

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func nums2pass(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + 'a' - 1
	}
	return string(bytes)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	if s == "a" || s == "zzzzzzzzzzzzzzzzzzzz" {
		puts("NO")
		return
	}

	nums := make([]int, n)
	for i := range nums {
		nums[i] = int(s[i]-'a') + 1
	}
	sort.Sort(sort.IntSlice(nums))

	var pass string
	hash := sum(nums...)
	if nums[n-1] == num('a') {
		// 2文字以上のaだけの列
		pass = string('a' + n - 1)
	} else if nums[0] == num('z') || n == 1 {
		// 19文字以下のzだけの列
		// もしくは"a"以外の1文字の列
		nums[0]--
		nums = append(nums, num('a'))
		pass = nums2pass(nums)
	} else if n == 20 {
		// "z"*20以外の20文字の列
		pass = strings.Repeat("z", hash/num('z'))
		if r := hash % num('z'); r > 0 {
			pass += string('a' + r - 1)
		}
		// 1文字だけzと異なる文字，それ以外はzの列に起きうる
		if pass == s {
			pass = pass[1:] + string(pass[0])
		}
	} else {
		nums[0]++
		nums[n-1]--
		pass = nums2pass(nums)
	}
	puts(pass)
}
