// Description: Fibonacci sequence implementation using different approaches.
package fibonacci

func FibR(n int) int64 {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  }

  return FibR(n-1) + FibR(n-2)
}

func FibC(n int, f []int64) int64 {
  if f[n] != -1 {
    return f[n]
  }
  f[n] = FibC(n-1, f) + FibC(n-2, f)
  return f[n]
}

func FibCDriver(n int) int64 {
  f := make([]int64, n+1)
  f[0] = 0
  f[1] = 1
  for i := 2; i <= n; i++ {
    f[i] = -1 
  }

  return FibC(n, f)
}

func FibDP(n int) int64 {
  f := make([]int64, n+1)
  f[0] = 0
  f[1] = 1
  
  for i := 2; i <= n; i++ {
    f[i] = f[i-1] + f[i-2]
  }

  return f[n]
}

func FibUltimate(n int) int64 {
  var back1, back2 int64 = 1, 0
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
