package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	ans := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if a <= sum && sum <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
