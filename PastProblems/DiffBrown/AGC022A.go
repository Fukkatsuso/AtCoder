package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	// table[i][j] = i文字目より前に文字jが存在するか
	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, 26)
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			for j := 0; j < 26; j++ {
				table[i][j] = table[i-1][j]
			}
		}
		table[i][s[i]-'a'] = true
	}

	ans := []byte(s)
	if len(ans) == 26 {
		// 末尾からみて辞書順最小の次の文字列を決定
		for loop := true; loop && len(ans) > 0; {
			last := len(ans) - 1
			for i := ans[last] - 'a' + 1; i <= 'z'-'a'; i++ {
				if !table[last][i] {
					ans[last] = 'a' + i
					loop = false
					break
				}
			}
			if loop {
				ans = ans[:last]
			}
		}
	} else {
		// 末尾に未出の文字を追加
		for i := 0; ; i++ {
			if !table[n-1][i] {
				ans = append(ans, byte('a'+i))
				break
			}
		}
	}

	if len(ans) > 0 {
		puts(string(ans))
	} else {
		puts(-1)
	}
}

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
