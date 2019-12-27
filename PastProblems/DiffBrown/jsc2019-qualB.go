// 答え見た

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func inputIntSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	a := inputIntSlice(n)

	cnt := []int{0, 0}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i] > a[j] {
				// あるAiの内部で発生するもの
				if i < j {
					cnt[0]++
				}
				// Ai, Ajの間で発生するもの
				cnt[1]++
			}
		}
	}

	mod := int64(1000000007)
	sel := int64(k*(k-1)) / 2 % mod
	ans := ((int64(cnt[0])*int64(k))%mod + (int64(cnt[1])*sel)%mod) % mod
	fmt.Println(ans)
}
