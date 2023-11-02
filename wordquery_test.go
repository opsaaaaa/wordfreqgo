package main

import (
  "testing"
)


func Test_WordQuery_TopN(t *testing.T) {
  wq := NewWordQuery("en")
  result, err := wq.TopN(10)
  assetErrNil(t, err)
  asertSame(t, "the", result[0])
  asertSame(t, 10, len(result))
}

func Test_WordQuery_Lookup(t *testing.T) {
  wq := NewWordQuery("en")
  wq.unitConversion = CbToCb
  results, err := wq.Lookup("the")
  assetErrNil(t, err)
  asertSame(t, results[0].word, "the")
  asertSame(t, results[0].val, 127.0)
}

func Benchmark_WordQuery_TopN(b *testing.B) {
  wq := NewWordQuery("en")

  wq.size = "large"
  b.Run("large 1000", func(b *testing.B) { benchmark__WordQuery_TopN(b, wq, 1000) })
  b.Run("large 10000", func(b *testing.B) { benchmark__WordQuery_TopN(b, wq, 10000) })

  wq.size = "small"
  b.Run("small 1000", func(b *testing.B) { benchmark__WordQuery_TopN(b, wq, 1000) })
  b.Run("small 10000", func(b *testing.B) { benchmark__WordQuery_TopN(b, wq, 10000) })
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

func benchmark__WordQuery_TopN(b *testing.B, wq *WordQuery, top int) {
  for n := 0; n < b.N; n++ {
    _, _ = wq.TopN( top )
  }
}

func benchmark_WordQuery_Lookup(b *testing.B, wq *WordQuery, input ...string) {
  for n := 0; n < b.N; n++ {
    _, _ = wq.Lookup( input... )
  }
}
