package main

import (
	"fmt"
  "math"
)

func Sqrt(x float64) float64 {
  s := float64(0)
  for z := float64(1); math.Abs(s - z) > 1e-15; z = z - (z*z - x)/(2*z) {
      s = z
  }
  return s
}

func main() {
	fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}

// $ go run flowcontrol8.go
// 1.4142135623730951
// 1.4142135623730951
