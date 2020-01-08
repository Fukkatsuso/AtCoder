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
	a := make([][]int, 4)
	for i := range a {
		a[i] = nextInts(4)
	}
	can := false
	// right
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == a[i][j+1] {
				can = true
			}
		}
	}
	// left
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if a[i][j] == a[i][j-1] {
				can = true
			}
		}
	}
	// up
	for i := 3; i > 0; i-- {
		for j := 0; j < 4; j++ {
			if a[i][j] == a[i-1][j] {
				can = true
			}
		}
	}
	// down
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if a[i][j] == a[i+1][j] {
				can = true
			}
		}
	}

	if can {
		fmt.Println("CONTINUE")
	} else {
		fmt.Println("GAMEOVER")
	}
}
