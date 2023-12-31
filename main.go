package main

import (
	"fmt"
	"os"
)


func main() {
  wq := NewWordQuery("en")
  results, err := wq.Lookup( os.Args[1:]... )
  if err != nil { panic(err) }
  for _, wci := range results {
    fmt.Printf("%v: %v\n", wci.word, wci.val)
  }
}


