package main

import "math"

func CbToFq(cb float64) float64 {
  return math.Pow(10.00, -math.Abs(cb)/100.00)
}
func CbToFpmw(cb float64) float64 {
  return math.Pow(10.00, -math.Abs(cb)/100.00) * 1000000
}
func CbToFpbw(cb float64) float64 {
  return math.Pow(10.00, -math.Abs(cb)/100.00) * 1000000000
}
func CbToZipf(cb float64) float64 {
  return (-math.Abs(cb) + 900.00) / 100.00
}


func FqToFpmw(fq float64) float64 {
	return fq * 1000000
}
func FqToFpbw(fq float64) float64 {
	return fq * 1000000000
}
func FqToZipf(fq float64) float64 {
	return math.Log10(fq) + 9
}
func FqToCb(fq float64) float64 {
	return math.Abs(math.Log10(fq) * 100.0)
}


/*
  HalfHarmonicMean is an average funciton.
  but it is kindof incorrect in this case.

  the phrase "the duke" can't be more common than duke.
  with this mean function get a sortof averadge.
  It works bettwe with "new york"
  but anything with common words, like "the",
  will end up being distorted

  the duke: 198 duke: 455 the: 127
  new york: 312 new: 275 york: 363
  it is: 198 it: 205 is: 193

  my gut is saying it should be lowest value word plus some avaradge.
  Or you add them all up, and then divid the differnece between the lowest and the sum.

  I looked up adding probabolities.
  P(A or B) = P(A) + P(B) - P(A and B)

  P(A or B): The probability that either event A or event B (or both) will occur.
  P(A), P(B): Tthe probability of event A/B occurring on its own.
  P(A and B): The probability that both events A and B occur simultaneously.

  we're looking for P(A and B).

  P(A and B) ≈ P(A) * P(B) // assuming that A and B are independent

  that should work if p(A/B) are both proportions between 1 and 0
  But I suspect that we'd need to convert them away from cb first.

  The other thought occurd to me. what if we need to convert it to a proportion for half harmonic mean to work properly?
*/ 

func HalfHarmonicMeanArr(args []int) int {
  sum := 0.0
  for _, arg := range args {
    sum += 1/float64(arg)
  }
  if sum == 0.0 {return int(sum)}
  return int( float64(len(args))/sum )
}

func HalfHarmonicMean(a, b float64) float64 {
    if a == 0 || b == 0 {
        return 0
    }
    return 2/(1/a + 1/b)
}










