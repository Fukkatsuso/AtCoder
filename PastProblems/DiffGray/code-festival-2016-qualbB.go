// コメントアウトの部分で入力受け取るとREになる

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	// sc.Split(bufio.ScanWords)
	// n, a, b := nextInt(), nextInt(), nextInt()
	// s := nextMega()
	var n, a, b int
	var s string
	fmt.Scan(&n, &a, &b, &s)
	rank := 0
	sum := 0
	for i := 0; i < n; i++ {
		ok := false
		if s[i] == 'a' && sum < a+b {
			sum++
			ok = true
		} else if s[i] == 'b' {
			rank++
			if sum < a+b && rank <= b {
				ok = true
				sum++
			}
		}
		if ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
