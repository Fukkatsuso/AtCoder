package main

import (
	"fmt"
)

func printX(b byte, x int) {
	for i := 0; i < x; i++ {
		fmt.Printf("%s", string(b))
	}
}

func main() {
	var sx, sy int
	var tx, ty int
	fmt.Scan(&sx, &sy, &tx, &ty)
	dx := tx - sx
	dy := ty - sy

	for i := 0; i < 2; i++ {
		printX('L', i)
		printX('U', dy+i)
		printX('R', dx+i)
		printX('D', i)
		printX('R', i)
		printX('D', dy+i)
		printX('L', dx+i)
		printX('U', i)
	}

	fmt.Printf("\n")
}
