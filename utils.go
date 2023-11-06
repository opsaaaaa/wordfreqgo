package main

// import "math"

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

func UniqStr(arr []string) []string {
  visited := make(map[string]bool)
  result := make([]string,0)
  for _, e := range arr {
    if _, ok := visited[e]; !ok {
      result = append(result, e)
      visited[e] = true
    }
  }
  return result
}

// func getValNestedMap(m map[string]interface{}) interface{} {
//     for i := range m {
//         nestedMap, ok := m[i].(map[string]interface{})
//         if ok {
//             return getValNestedMap(nestedMap)
//         }
//         return m[i]
//     }

//     return nil
// }


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




