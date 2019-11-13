package main

import "fmt"

func Bills(n, y int) (int, int, int) {
	if y < 1000*n || 10000*n < y {
		return -1, -1, -1
	}

	for i := 0; i <= y/10000; i++ {
		for j := 0; j <= (y-10000*i)/5000; j++ {
			k := (y - 10000*i - 5000*j) / 1000
			if i+j+k == n {
				return i, j, k
			}
		}
	}
	return -1, -1, -1
}

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	fmt.Println(Bills(n, y))
}
