package main

import "math"

/*
  | En Words     | Zipf   | Cb (abs) | Percentage % | Fq Propotion  | Fpmw      | Fpbw        |
  |--------------|--------|----------|--------------|---------------|-----------|-------------|
  |              | 9.0    | 0        | 100%         | 1             | 1000000   | 1000000000  |
  | the          | 8.0    | 100      | 10%          | 0.1           | 100000    | 100000000   |
  | that, for    | 7.0    | 200      | 1%           | 0.01          | 10000     | 10000000    |
  | said, way    | 6.0    | 300      | 0.1%         | 0.001         | 1000      | 1000000     |
  | radio, plans | 5.0    | 400      | 0.01%        | 0.0001        | 100       | 100000      |
  | prizes, bail | 3.0    | 500      | 0.001%       | 0.00001       | 10        | 10000       |
  | sparing      | 2.0    | 600      | 0.0001%      | 0.000001      | 1         | 1000        |
  | cryptology   | 1.0    | 700      | 0.00001%     | 0.0000001     | 0.1       | 100         |
  | microcapsule | 0      | 800      | 0.000001%    | 0.00000001    | 0.01      | 10          |
  |              | -1     | 900      | 0.0000001%   | 0.000000001   | 0.001     | 1           |
  |              | -2     | 1000     | 0.00000001%  | 0.0000000001  | 0.0001    | 0.1         |
  |              | -3     | 1100     | 0.000000001% | 0.00000000001 | 0.00001   | 0.01        |
  |              | ∞      | ∞        | 0%           | 0             | 0         | 0           |
*/


/*
  Cb is a word frequency from of logarithmic centibel scale.
  It is the word frequency unit used the dataset from the python wordfq program.
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
func CbToCb(cb float64) float64 {
  return math.Abs(cb)
}

/*
  ZipF is log10 of frequency per billion words
  Named after the American linguist George Kingsley Zipf

  Advantages
  - Its human readable and its a known common standerd.

  Disavantages
  - It requires decimials for accuracy.
  - Technically it can cross 0 with extremely rare items in large datasets.
*/
func ZipfToFq(zipf float64) float64 {
  return math.Pow(10.00, zipf) / 1e9
}
func ZipfToFpmw(zipf float64) float64 {
  return math.Pow(10.00, zipf) / 1000
}
func ZipfToFpbw(zipf float64) float64 {
  return math.Pow(10.00, zipf)
}
func ZipfToCb(zipf float64) float64 {
  return math.Abs((zipf * 100.0) - 900.00)
}


/*
  Fq frequency represented as a proportion between 0 and 1
  occurances count for a word in the corpus divided by total words in corpus
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
  fpmw frequency per million words.
  or the number of times a word occurs in one million words

  a fpmw of 1 means that word occurs 1 once on average for every million words
  a fpmw of 1,000,000 would mean every word/token in the corpus was the same.

  Advantages
  - Its straight forward to calculated and understand.
  - corpus size doesn't change the relative value.
  - Its a old standard.

  Disadvantages:
  - the issue with fpmw is that rare words can have a fpmw of less than 1
  - and its not easy for humans to compare.
*/
// func FpmwToFq(fpmw float64) float64 {
//  return fpmw / 1000000.0
// }
// func FpmwToFpbw(fpmw float64) float64 {
//  return fpmw * 1000
// }
// func FpmwToZipf(fpmw float64) float64 {
//  return math.Log10(fpmw * 1000)
// }
func FpmwToCb(fpmw float64) float64 {
 return math.Abs(math.Log10(fpmw / 1000000.0) * 100.0)
}
// func CalcFpmw(occurances float64, total float64) float64 {
//  return (occurances / total) * 1000000
// }


/*
  fpbw frequency per billion words.
  the same as fpmw but with a billion instead of million.

  The advantages over fpbw is that values are far less likely to dip below 1
*/
// func FpbwToFq(fpbw float64) float64 {
//   return fpbw / 1000000000.0
// }
// func FpbwToFpmw(fpbw float64) float64 {
//   return fpbw / 1000
// }
// func FpbwToZipf(fpbw float64) float64 {
//   return math.Log10(fpbw)
// }
func FpbwToCb(fpbw float64) float64 {
  return math.Abs(math.Log10(fpbw / 1000000000.0) * 100.0)
}
// func CalcFpbw(occurances float64, total float64) float64 {
//   return (occurances / total) * 1000000000
// }
