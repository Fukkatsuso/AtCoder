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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	// table[i]: iで何回割り切れるか
	table := make([]int, n+1)

	for i := 2; i <= n; i++ {
		num := i
		for j := 2; j <= n; j++ {
			for num%j == 0 {
				num /= j
				table[j]++
			}
		}
	}

	ans := 1
	for _, v := range table {
		if v == 0 {
			continue
		}
		ans = (ans * (v + 1)) % 1000000007
	}
	fmt.Println(ans)
}
