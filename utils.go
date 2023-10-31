package main

import "math"

func CbToFq(cb int) float64 {
  return math.Pow(10.00, -float64(cb)/100.00)
}

func CbToFpmw(cb int) float64 {
  return math.Pow(10.00, -float64(cb)/100.00) * 1000000
}

func CbToFpbw(cb int) float64 {
  return math.Pow(10.00, -float64(cb)/100.00) * 1000000000
}

func CbToZipf(cb int) float64 {
  return (-float64(cb) + 900.00) / 100.00
}



// HalfHarmonicMean is an average funciton.
// but it is kindof incorrect in this case.

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








