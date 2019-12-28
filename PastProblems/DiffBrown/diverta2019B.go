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
	r, g, b, n := nextInt(), nextInt(), nextInt(), nextInt()
	ans := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if n >= (r*i+g*j) && (n-(r*i+g*j))%b == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
