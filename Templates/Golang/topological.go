package main

// トポロジカルソート
// n: グラフgのノード数
// g: グラフ
// indegree: 各ノードの入次数
func topologicalSort(n int, g *Graph, indegree []int) []int {
	sorted := make([]int, 0)

	que := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		for _, e := range g.Edges[v] {
			u := e.to
			indegree[u]--
			if indegree[u] == 0 {
				que = append(que, u)
			}
		}
		sorted = append(sorted, v)
	}

	return sorted
}
