// Description: This package provides a function to calculate the binomial coefficient of n and k.
package binomialCoefficients

func BinomialCoefficient(n, k int) int64 {
  bc := make([][]int64, n+1)
  for i := range bc {
    bc[i] = make([]int64, n+1)
  }

  // Base cases
  for i := 0; i <= n; i++ {
    bc[i][0] = 1
    bc[i][i] = 1
  }

  // Recurrence representation
  for i := 2; i <= n; i++ {
    for j := 1; j < i; j++ {
      bc[i][j] = bc[i-1][j-1] + bc[i-1][j]
    }
  }

  return bc[n][k]
}
