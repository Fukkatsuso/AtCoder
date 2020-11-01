package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()

	// 2桁以下のsについて
	if len(s) == 1 {
		if s == "8" {
			puts("Yes")
			return
		}
	} else if len(s) == 2 {
		x, y := int(s[0]-'0'), int(s[1]-'0')
		if (x*10+y)%8 == 0 || (y*10+x)%8 == 0 {
			puts("Yes")
			return
		}
	}
	if len(s) < 3 {
		puts("No")
		return
	}

	cnt := make([]int, 10)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}

	// xを作れるか?
	ok := false
	for x := 112; x < 1000; x += 8 {
		c := map[int]int{}
		for y := 1; y <= 100; y *= 10 {
			c[(x/y)%10]++
		}
		ok2 := true
		for k, v := range c {
			if len(s) < 3 && k == 0 {
				continue
			}
			ok2 = ok2 && (cnt[k] >= v)
		}
		ok = ok || ok2
	}

	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
}
