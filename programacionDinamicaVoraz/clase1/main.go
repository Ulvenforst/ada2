// clase 1: Introducción a Programación Dinámica, caso: secuencia Fibonacci.
package main

import (
	"fmt"
	"time"

	"github.com/Ulvenforst/ada2/programacionDinamicaVoraz/clase1/binomialCoefficients"
	"github.com/Ulvenforst/ada2/programacionDinamicaVoraz/clase1/fibonacci"
)

func main() {
  start := time.Now()
  // var fibNumber int64 = fibonacci.FibR(4)
  // var fibNumber int64 = fibonacci.FibCDriver(50)
  // var fibNumber int64 = fibonacci.FibDp(50)
  var fibNumber int64 = fibonacci.FibUltimate(50)
  var binomial int64 = binomialCoefficients.BinomialCoefficient(5,2)
  
  fmt.Println(fibNumber, "Time: ", time.Since(start))
  fmt.Println("binomial: ", binomial)
}
