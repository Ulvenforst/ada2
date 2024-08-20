package main

import (
	"fmt"
	"github.com/Ulvenforst/ada2/programacionDinamicaVoraz/clase2/leastCommonSubsequence"
)

func main() {
	X := "BCBDAB"
	Y := "BDCBA"
	m := len(X)
	n := len(Y)

	c, b := leastCommonSubsequence.LCSLength(X, Y, m, n)

	fmt.Printf("Length of LCS: %d\n", c[m][n])
	fmt.Print("LCS: ")
	leastCommonSubsequence.PrintLCS(b, X, m, n)
	fmt.Println()
}
