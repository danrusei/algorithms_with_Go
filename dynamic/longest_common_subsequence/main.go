package main

import "fmt"

func lcs(s1 string, s2 string) int {

	m := len(s1)
	n := len(s2)

	//create the m x n matrix
	mat := make([][]int, m+1)
	for i := range mat {
		mat[i] = make([]int, n+1)
	}

	//Following steps build mat[m+1][n+1] in bottom up fashion.
	//Note that mat[i][j] contains length of LCS of s1[0..i-1] and s2[0..j-1]
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				mat[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				mat[i][j] = mat[i-1][j-1] + 1
			} else {
				mat[i][j] = max(mat[i-1][j], mat[i][j-1])
			}
		}
	}
	return mat[m][n]
}

func max(m int, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	fmt.Println(lcs("ABCDGH", "AEDFHR"))
	fmt.Println(lcs("AGGTAB", "GXTXAYB"))
}
