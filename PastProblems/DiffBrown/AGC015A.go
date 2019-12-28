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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	n, a, b := nextInt(), nextInt(), nextInt()
	if a > b {
		fmt.Println(0)
		return
	}
	if n-2 < 0 {
		if a == b {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	ans := int64(n-2)*(int64(b-a)) + 1
	fmt.Println(ans)
}
