// Description: Fibonacci sequence implementation using different approaches.
package fibonacci

const (
    UNKNOWN = -1
    MAXN    = 92
)

var f [MAXN + 1]int64

// Recursive Fibonacci
func FibR(n int) int64 {
  if n == 0 {
    return 0
  }

  if n == 1 {
    return 1
  }

  return FibR(n-1) + FibR(n-2)
}

// Caching Fibonacci
func fibC(n int) int64 {
  if f[n] == UNKNOWN {
    f[n] = fibC(n-1) + fibC(n-2)
  }

  return f[n]
}

func FibCDriver(n int) int64 {
  f[0] = 0
  f[1] = 1

  for i := 2; i <= n; i++ {
    f[i] = UNKNOWN
  }

  return fibC(n)
}

// Linear fibonacci
func FibDp(n int) int64 {
  f[0] = 0
  f[1] = 1

  for i := 2; i <= n; i++ {
    f[i] = f[i-1] + f[i-2]
  }

  return f[n]
}

// Ultimate version
func FibUltimate(n int) int64 {
  back2, back1 := int64(0), int64(1)
  var next int64

  if n == 0 {
    return 0
  }

  for i := 2; i < n; i++ {
    next = back1 + back2
    back2 = back1
    back1 = next
  }

  return back1 + back2
}

