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
	n := nextInt()
	a := inputIntSlice(n)

	cnt := make([]int, n)
	for i := range a {
		if a[i]%2 == n%2 {
			fmt.Println(0)
			return
		}
		cnt[a[i]]++
	}

	if n%2 == 0 && cnt[0] != 0 {
		fmt.Println(0)
		return
	} else if n%2 == 1 && cnt[0] != 1 {
		fmt.Println(0)
		return
	}
	for i := n - 1; i > 0; i -= 2 {
		if cnt[i] != 2 {
			fmt.Println(0)
			return
		}
	}

	ans := 1
	for i := 0; i < n/2; i++ {
		ans = (ans * 2) % 1000000007
	}
	fmt.Println(ans)
}
