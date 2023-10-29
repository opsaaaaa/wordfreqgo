
package main

import "testing"

func asertSame(t *testing.T, a, b interface{}) {
  if a != b {
    t.Errorf("Expected %v, but got %v", b, a)
  }
}

func Test_HalfHarmonicMean(t *testing.T) {
  asertSame(t, HalfHarmonicMean([]int{2, 4, 8, 16}), 5)
  asertSame(t, HalfHarmonicMean([]int{5}), 5)
}

