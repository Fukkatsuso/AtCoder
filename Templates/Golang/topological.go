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

// トポロジカルソートの組み合わせ
// n: ノード数
// to[i]: ノードiから他のノードに辺が伸びているかを{0,1}でbit管理
// O(n * 2^n)
// 全ノードの並べ方はdp[(1<<n)-1]で求まる
func topologicalSortPattern(n int, to []int) []int {
	dp := make([]int, 1<<n)
	dp[0] = 1
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 && i&to[j] == 0 {
				dp[i|(1<<j)] += dp[i]
			}
		}
	}
	return dp
}
