package main

import (
	"math"
	"testing"
)

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
  asertSame(t, HalfHarmonicMeanArr([]int{127, 400}), int(HalfHarmonicMean(127, 400)))
}

func Test_HalfHarmonicMean(t *testing.T) {
  asertSame(t, HalfHarmonicMean(275, 363), 469)
}


func Test_SomeMath(t *testing.T) {
  asertSame(t, CbToFq(0.0), 1.0)
  asertSame(t, CbToFq(100), 0.1)
  asertSame(t, FqToCb(1.0), 0.0)
  asertSame(t, math.Round(FqToCb(0.1)), 100.0)

  asertSame(t, CbToFq(100.0) * CbToFq(200.0), 0.001)
  asertSame(t,
    math.Round(HalfHarmonicMean(205.0, 193.0)),
    math.Round(FqToCb( HalfHarmonicMean(CbToFq(205.0), CbToFq(193.0)))),
  )

  asertSame(t,
    math.Round(HalfHarmonicMean(205.0, 193.0)),
    math.Round(FqToCb( CbToFq(205.0) * CbToFq(193.0) )),
  )



  // 397.99999999999994
  //205 193
  // a * a / 
  // asertSame(t, a, c)

}
