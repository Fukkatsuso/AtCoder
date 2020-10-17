// 写経AC
// https://www.slideshare.net/chokudai/code-formula2014-qualb
// https://at274.hatenablog.com/entry/2020/06/14/154107

package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := next(), next()
	n := len(a)

	diff := make([]int, 0)
	dup := -1
	appeared := map[byte]bool{}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			diff = append(diff, i)
		}
		if appeared[a[i]] {
			dup = i
		}
		appeared[a[i]] = true
	}
	// aとbの違いが6箇所に満たない場合，複数回出現する文字(のインデックス)を追加
	if dup >= 0 {
		for i := 0; i < 6-len(diff); i++ {
			diff = append(diff, dup)
		}
	}

	// aとbの違いが6箇所以上ならNO
	if len(diff) > 6 {
		puts("NO")
		return
	}

	check := func(s string, cnt int) bool {
		return (dup >= 0 || cnt == 3) && (s == b)
	}

	ok := false
	var dfs func(s []byte, depth int)
	dfs = func(s []byte, depth int) {
		if depth == 3 {
			return
		}
		for i, x := range diff {
			for j, y := range diff {
				if i >= j {
					continue
				}
				s[x], s[y] = s[y], s[x]
				ok = ok || check(string(s), depth+1)
				dfs(s, depth+1)
				s[x], s[y] = s[y], s[x]
			}
		}
	}

	dfs([]byte(a), 0)
	if ok {
		puts("YES")
	} else {
		puts("NO")
	}
}
