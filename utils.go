package main

import "math"

/*
  Cb is a word frequency from of logarithmic centibel scale.

  practical range -127(the) to -799
  actuall range is 0 to -900(or less)

  Cb is the word frequency unit used the dataset from the python wordfq program.
  https://github.com/rspeer/wordfq

  > 0 cB represents a word that occurs with probability 1, so it is the only
  > word in the data (this of course doesn't happen). -200 cB represents a
  > word that occurs once per 100 tokens, -300 cB represents a word that
  > occurs once per 1000 tokens, and so on.

  Advantages
  - Its very similar to zipf, but with a different scale and 0 point.
  - Its really good for storage sizes.
  - Its always less than 0, so rare values cant cross 0.
  - and numbers are larger, so you dont need decimils for reasonable accuracy.
  - you can easilly save them as positive integers.

  Disadvantages
  - its less human readable.

  In the wordfq program they 'bin' the data to reduce the file size further.
  array[ bin[ "words", ...], ... ]
  The index of the bin represents the positive frequency value.
  you end up with a lot of leading empty bins, but after that it gets really effecient.
  I've decided to use .tsv and line/row numbers instead.
  either way its quite compact.
*/

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

/*
  Fq frequency represented as a proportion between 0 and 1
  occurances count for a word in the corpus divided by total words in corpus

  practicle range 0.053(the) 0.00000001(trella)
  actual range 0 to 1
*/
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
  Combind the probability of both A and B and C occuring, in the CB format.
  P(A and B) ≈ P(A) * P(B)
  assuming that A and B are independent
*/
func CbAndProbabilities(args ...int) float64 {
  sum := 1.0
  for _, arg := range args {
    sum = sum * CbToFq(float64(arg))
  }
  return FqToCb(sum)
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










