
package main

import "testing"

func asertSame(t *testing.T, a, b interface{}) {
  if a != b {
    t.Errorf("Expected %v, but got %v", b, a)
  }
}

func assetErrNil(t *testing.T, err error) {
  if err != nil {
    t.Errorf("Expected error to be nil.\n%v", err)
  }
}

func Test_HalfHarmonicMeanArr(t *testing.T) {
  asertSame(t, HalfHarmonicMeanArr([]int{2, 4, 8, 16}), 5)
  asertSame(t, HalfHarmonicMeanArr([]int{5}), 5)
  // asertSame(t, HalfHarmonicMeanArr([]int{127, 400}), int(HalfHarmonicMean(-127, -400)))
}


/*
  the potato
  127 489
  a+b/2 wrong

  a+b 2
*/

func Test_SomeMath(t *testing.T) {
  asertSame(t, 1, 1)

}
