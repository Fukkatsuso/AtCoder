// AC
package main

import (
	"fmt"
)

type Edge struct {
	to int
	id int
}

func dfs(edges [][]Edge, visited []bool, colors []int, from, now int) {
	if visited[now] {
		return
	}
	visited[now] = true

	ng := -1
	for _, e := range edges[now] {
		if e.to == from {
			ng = colors[e.id]
		}
	}

	color := 1
	for _, e := range edges[now] {
		if color == ng {
			color++
		}
		if e.to == from || visited[e.to] || colors[e.id] != 0 {
			continue
		}
		colors[e.id] = color
		color++
		dfs(edges, visited, colors, now, e.to)
	}
}

func max(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	edges := make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		edges[a] = append(edges[a], Edge{to: b, id: i})
		edges[b] = append(edges[b], Edge{to: a, id: i})
	}

	colors := make([]int, n-1)
	dfs(edges, make([]bool, n), colors, -1, 0)

	fmt.Println(max(colors))
	for _, col := range colors {
		fmt.Println(col)
	}
}
