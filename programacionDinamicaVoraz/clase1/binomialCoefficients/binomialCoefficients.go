// Description: This package provides a function to calculate the binomial coefficient of n and k.
package binomialCoefficients

const MAXN = 100

func BinomialCoefficient(n, k int) int64 {
  var bc [MAXN+1][MAXN+1]int64

  for i := 0; i <= n; i++ {
    bc[i][0] = 1
  }

  for j := 0; j <= n; j++ {
    bc[j][j] = 1
  }

  for i := 2; i <= n; i++ {
    for j := 1; j < i; j++ {
      bc[i][j] = bc[i-1][j-1] + bc[i-1][j]
    }
  }

  return bc[n][k]
}
