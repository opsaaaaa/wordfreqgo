package main

import "math"

func HalfHarmonicMeanArr(args []int) int {
  if len(args) == 1 {
    return args[0]
  }
  if len(args) == 0 {
    return -1
  }

  sumReciprocals := float64(0)
  for _, arg := range args {
    sumReciprocals += 1.0 / float64(arg)
  }
  return int( math.Round( float64(len(args) + 1) / sumReciprocals))
}

func HalfHarmonicMean(a, b float64) float64 {
    if a == 0 || b == 0 {
        return 0
    }
    return 2/(1/a + 1/b)
}








