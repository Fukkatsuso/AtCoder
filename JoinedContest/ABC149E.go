// WA

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	a := nextInts(n)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	ans := int64(0)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			ans += int64(a[i] + a[j])
			cnt++
			if cnt == m {
				break
			}
			if i == j {
				continue
			}
			ans += int64(a[i] + a[j])
			cnt++
			if cnt == m {
				break
			}
		}
		if cnt == m {
			break
		}
	}
	fmt.Println(ans)
}
