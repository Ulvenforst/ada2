package matrixmultiplication

func RectangularMatrixMultiply(A, B, C [][]float64, p, q, r int) {
    for i := 0; i < p; i++ {
        for j := 0; j < r; j++ {
            for k := 0; k < q; k++ {
                C[i][j] += A[i][k] * B[k][j]
            }
        }
    }
}
