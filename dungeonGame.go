package leetcodedungeongame

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	m++
	n++

	healthNeeded := make([][]int, m)

	for i := 0; i < m; i++ {
		healthNeeded[i] = make([]int, n)
		healthNeeded[i][n-1] = 1e9
	}

	for i := 0; i < n; i++ {
		healthNeeded[m-1][i] = 1e9
	}

	healthNeeded[m-2][n-1], healthNeeded[m-1][n-2] = 1, 1

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			healthNeeded[i][j] = min(healthNeeded[i+1][j], healthNeeded[i][j+1]) - dungeon[i][j]
			if healthNeeded[i][j] <= 0 {
				healthNeeded[i][j] = 1
			}
		}
	}

	return healthNeeded[0][0]
}

func min(x, y int) int {
	ans := x
	if y < x {
		ans = y
	}

	return ans
}
