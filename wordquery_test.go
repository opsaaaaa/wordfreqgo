package main

import (
  "testing"
)

func Test_WordQuery_Lookup(t *testing.T) {
  wq := NewWordQuery("en")
  wq.unitConversion = CbToCb
  results, err := wq.Lookup("the")
  assetErrNil(t, err)
  for k, v := range results {
    asertSame(t, k, "the")
    asertSame(t, v, 127.0)
  }
}


func benchmark_WordQuery_Lookup(b *testing.B, wq *WordQuery, input ...string) {
  for n := 0; n < b.N; n++ {
    _, _ = wq.Lookup( input... )
  }
}

func Benchmark_WordQuery_Lookup(b *testing.B) {
  wq := NewWordQuery("en")

  wq.size = "large"
  b.Run("worse case large", func(b *testing.B) {  benchmark_WordQuery_Lookup(b, wq, "thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase") } )
  b.Run("best case large", func(b *testing.B) {  benchmark_WordQuery_Lookup(b, wq, "the") } )

  wq.size = "small"
  b.Run("worse case small", func(b *testing.B) {  benchmark_WordQuery_Lookup(b, wq, "thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase") } )
  b.Run("best case small", func(b *testing.B) {  benchmark_WordQuery_Lookup(b, wq, "the") } )
}


