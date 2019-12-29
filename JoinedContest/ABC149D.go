// 謎のRE

package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	r, s, p := nextInt(), nextInt(), nextInt()
	t := next()

	ans := int64(0)
	win := make([]bool, n)
	for i := 0; i < n; i++ {
		point := 0
		switch t[i] {
		case 'r':
			point = p
		case 's':
			point = r
		case 'p':
			point = s
		default:
			point = 0
		}
		if (i >= k) && (t[i-k] == t[i]) && win[i-k] {
			point = 0
			win[i] = false
		} else {
			win[i] = true
		}
		ans += int64(point)
	}
	fmt.Println(ans)
}
