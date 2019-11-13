package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n+1)
	x := make([]int, n+1)
	y := make([]int, n+1)
	t[0] = 0
	x[0] = 0
	y[0] = 0
	for i := 1; i < n+1; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}

	can := true

	for i := 0; i < n; i++ {
		dt := t[i+1] - t[i]
		dist := (x[i+1] - x[i]) + (y[i+1] - y[i])
		if dt < dist {
			can = false
			break
		}
		if dif := dt - dist; dif%2 == 1 {
			can = false
			break
		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
