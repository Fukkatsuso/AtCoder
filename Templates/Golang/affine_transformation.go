package main

// 変換行列
var (
	// 反時計回り(90度)
	rot90 = func() [][]int {
		return [][]int{
			{0, -1, 0},
			{1, 0, 0},
			{0, 0, 1},
		}
	}
	// 時計回り(90度)
	rotRev90 = func() [][]int {
		return [][]int{
			{0, 1, 0},
			{-1, 0, 0},
			{0, 0, 1},
		}
	}

	// x=pについて線対称
	symmX = func(p int) [][]int {
		return [][]int{
			{-1, 0, 2 * p},
			{0, 1, 0},
			{0, 0, 1},
		}
	}

	// y=pについて線対称
	symmY = func(p int) [][]int {
		return [][]int{
			{1, 0, 0},
			{0, -1, 2 * p},
			{0, 0, 1},
		}
	}
)

// 行列m1,m2の積
func matixProduct(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		res[i] = make([]int, len(m2[i]))
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := range m1[i] {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

// アフィン変換(2次元)
func affineTransformation(x, y int, mat [][]int) (int, int) {
	res := matixProduct(mat, [][]int{{x}, {y}, {1}})
	x, y = res[0][0], res[1][0]
	return x, y
}
