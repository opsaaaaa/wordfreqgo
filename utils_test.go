package main

import (
	"testing"
)

func asertSame(t *testing.T, a, b interface{}) {
  if a != b {
    t.Errorf("Expected %v, but got %v", a, b)
  }
}

func assetErrNil(t *testing.T, err error) {
  if err != nil {
    t.Errorf("Expected error to be nil.\n%v", err)
  }
}
// asset assert asert
func assertStringSliceSame(t *testing.T, a, b []string) {
  if len(a) != len(b) {
    t.Errorf("Expected %v, but got %v", a, b)
    return
  }
  for i := range a {
    if a[i] != b[i] {
      t.Errorf("Expected %v, but got %v", a, b)
      return
    }
  }
}

func Test_Uniq(t *testing.T) {
  expected := []string{"one", "two", "three"}
  actuall := UniqStr([]string{"one","one","two","one","three","two"})
  assertStringSliceSame(t, expected, actuall)
}

// func Test_SomeMath(t *testing.T) {
//   asertSame(t, CbToFq(0.0), 1.0)
//   asertSame(t, CbToFq(100), 0.1)
//   asertSame(t, FqToCb(1.0), 0.0)
//   asertSame(t, math.Round(FqToCb(0.1)), 100.0)

//   asertSame(t, CbToFq(100.0) * CbToFq(200.0), 0.001)
//   asertSame(t,
//     math.Round(HalfHarmonicMean(205.0, 193.0)),
//     math.Round(FqToCb( HalfHarmonicMean(CbToFq(205.0), CbToFq(193.0)))),
//   )

//   asertSame(t,
//     math.Round(HalfHarmonicMean(205.0, 193.0)),
//     math.Round(FqToCb( CbToFq(205.0) * CbToFq(193.0) )),
//   )

// }
