package leastCommonSubsequence

import "fmt"

func LCSLength(X, Y string, m, n int) ([][]int, [][]string) {
	c := make([][]int, m+1)
	b := make([][]string, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
		b[i] = make([]string, n+1)
	}

	for i := 1; i <= m; i++ {
		c[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		c[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = "↖"
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = "↑"
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = "←"
			}
		}
	}

	return c, b
}

func PrintLCS(b [][]string, X string, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == "↖" {
		PrintLCS(b, X, i-1, j-1)
		fmt.Print(string(X[i-1]))
	} else if b[i][j] == "↑" {
		PrintLCS(b, X, i-1, j)
	} else {
		PrintLCS(b, X, i, j-1)
	}
}
