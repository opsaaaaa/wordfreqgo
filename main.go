package main

import (
	"fmt"
	"os"
)


func main() {
  wq := NewWordQuery("en")
  results, err := wq.LookupMultiple( os.Args[1:] )
  if err != nil { panic(err) }
  for q, r := range results {
    fmt.Printf("%v: %v\n", q, r)
  }
}


