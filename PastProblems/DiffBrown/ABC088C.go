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
	c := make([][]int, 3)
	for i := range c {
		c[i] = inputIntSlice(3)
	}

	ok := false
	for a1 := 0; a1 <= 100; a1++ {
		a := []int{a1, c[1][0] - c[0][0] + a1, c[2][0] - c[0][0] + a1}
		b := []int{c[0][0] - a1, c[0][1] - a1, c[0][2] - a1}
		if a[1] < 0 || a[2] < 0 {
			continue
		}
		if b[0] < 0 || b[1] < 0 || b[2] < 0 {
			continue
		}
		ok = true
		for i := 1; i < 3; i++ {
			for j := 1; j < 3; j++ {
				ok = ok && (a[i]+b[j] == c[i][j])
			}
		}
		if ok {
			break
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
