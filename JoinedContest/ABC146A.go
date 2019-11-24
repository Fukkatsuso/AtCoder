package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var ans int
	switch s {
	case "SUN":
		ans = 7
	case "MON":
		ans = 6
	case "TUE":
		ans = 5
	case "WED":
		ans = 4
	case "THU":
		ans = 3
	case "FRI":
		ans = 2
	case "SAT":
		ans = 1
	}
	fmt.Println(ans)
}
