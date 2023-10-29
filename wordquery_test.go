package main

import (
  "testing"
)

func benchmark_WordQuery_LookupMultiple(b *testing.B, wq *WordQuery, input []string) {
  for n := 0; n < b.N; n++ {
    _, _ = wq.LookupMultiple( input )
  }
}


func Benchmark_WordQuery_LookupMultiple(b *testing.B) {
  wq := NewWordQuery("en")

  wq.size = "large"
  b.Run("worse case large", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, []string{"thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase"}) } )
  b.Run("best case large", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, []string{"the"}) } )

  wq.size = "small"
  b.Run("worse case small", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, []string{"thisistotallynotareallworldandsoweewillhaveto runallthewayto","theendofthefile worsecase"}) } )
  b.Run("best case small", func(b *testing.B) {  benchmark_WordQuery_LookupMultiple(b, wq, []string{"the"}) } )
}


