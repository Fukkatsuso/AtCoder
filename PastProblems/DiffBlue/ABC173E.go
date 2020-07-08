// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func modInv(a, mod int) int {
	b := mod
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		u -= t * v
		a, b = b, a
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

type Bigger []int

func (a Bigger) Len() int      { return len(a) }
func (a Bigger) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Bigger) Less(i, j int) bool {
	return abs(a[i]) > abs(a[j])
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n, k := nextInt(), nextInt()
	a := Bigger(nextInts(n))
	sort.Sort(a)

	ans := 1
	minus := 0
	for i := range a {
		if a[i] < 0 {
			minus++
		}
	}
	// 全て負の数かつKが奇数の場合
	// N==Kの場合
	if (minus == n && k%2 == 1) || n == k {
		// 小さいものをかけていく
		for i := n - k; i < n; i++ {
			ans *= (a[i] + mod) % mod
			ans %= mod
		}
		puts(ans)
		return
	}

	// 絶対値の大きい方からかけていく
	ans = 1
	minus = 0
	for i := 0; i < k; i++ {
		if a[i] < 0 {
			minus++
		}
		ans *= (a[i] + mod) % mod
		ans %= mod
	}
	// すでに正の数なら終了
	if minus%2 == 0 {
		puts(ans)
		return
	}

	// 取り除く候補
	// 0: 非負，1: 負
	rem := [2]int{}
	exist1 := [2]bool{}
	for i := k - 1; i >= 0; i-- {
		if exist1[0] && exist1[1] {
			break
		}
		if a[i] > 0 && !exist1[0] {
			rem[0] = a[i]
			exist1[0] = true
		} else if a[i] < 0 && !exist1[1] {
			rem[1] = a[i]
			exist1[1] = true
		}
	}
	// 追加する候補
	add := [2]int{}
	exist2 := [2]bool{}
	for i := k; i < n; i++ {
		if exist2[0] && exist2[1] {
			break
		}
		if a[i] > 0 && !exist2[0] {
			add[0] = a[i]
			exist2[0] = true
		} else if a[i] < 0 && !exist2[1] {
			add[1] = a[i]
			exist2[1] = true
		}
	}
	if exist1[0] && exist2[1] && exist1[1] && exist2[0] {
		if abs(add[1]*rem[1]) > abs(add[0]*rem[0]) {
			rem[0], add[1] = (rem[0]+mod)%mod, (add[1]+mod)%mod
			ans = (ans * modInv(rem[0], mod)) % mod
			ans = (ans * add[1]) % mod
		} else {
			rem[1], add[0] = (rem[1]+mod)%mod, (add[0]+mod)%mod
			ans = (ans * modInv(rem[1], mod)) % mod
			ans = (ans * add[0]) % mod
		}
	} else if exist1[0] && exist2[1] {
		rem[0], add[1] = (rem[0]+mod)%mod, (add[1]+mod)%mod
		// 非負の数を取り除き，負の数を加える
		ans = (ans * modInv(rem[0], mod)) % mod
		ans = (ans * add[1]) % mod
	} else if exist1[1] && exist2[0] {
		rem[1], add[0] = (rem[1]+mod)%mod, (add[0]+mod)%mod
		// 負の数を取り除き，非負の数を加える
		ans = (ans * modInv(rem[1], mod)) % mod
		ans = (ans * add[0]) % mod
	} else {
		ans = 0
	}

	puts(ans)
}
