package main

import (
  "testing"
)

func Test_WordQuery_Lookup(t *testing.T) {
  wq := NewWordQuery("en")
  fq, err := wq.Lookup("the")
  assetErrNil(t, err)
  asertSame(t, fq, 127)
}

func Test_WordQuery_LookupMultiple(t *testing.T) {
  wq := NewWordQuery("en")
  results, err := wq.LookupMultiple("the")
  assetErrNil(t, err)
  for k, v := range results {
    asertSame(t, k, "the")
    asertSame(t, v, 127)
  }
}


func benchmark_WordQuery_LookupMultiple(b *testing.B, wq *WordQuery, input ...string) {
  for n := 0; n < b.N; n++ {
    _, _ = wq.LookupMultiple( input... )
  }
}

func Benchmark_WordQuery_LookupMultiple(b *testing.B) {
  wq := NewWordQuery("en")

  wq.size = "large"
  b.Run("worse case large", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, "thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase") } )
  b.Run("best case large", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, "the") } )

  wq.size = "small"
  b.Run("worse case small", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, "thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase") } )
  b.Run("best case small", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, "the") } )
}


