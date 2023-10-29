package main

import (
	"fmt"
	"io"
	"strings"
)


const DATA_FILE_TSV_GZ = "data/%s_%s.tsv.gz"

type WordQuery struct {
  lang   string
  size   string
  max    int
  tokenize func(s string) []string
}

// Initialize WordQuery with default tokenize function
func NewWordQuery(lang string) *WordQuery {
  return &WordQuery{
    lang:     lang,
    size:     "small",
    max:      800,
    tokenize: strings.Fields,
  }
}

func (w *WordQuery) Lookup(query string) (int, error) {
  words := w.tokenize(query)
  found := len(words)
  results := make([]int, found)
  for i := range results {
    results[i] = w.max
  }

  err := WalkTsvGzCells(fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang), func(row, _ int, cell string) error {
    for i, word := range words {
      
      if cell == strings.ToLower(word) {
        results[i] = row
        found--
      }
    }

    if found == 0 { return io.EOF }
    return nil

  })
  return HalfHarmonicMean(results), err
}


func (w *WordQuery) LookupMultiple(queries []string) (map[string]int, error) {

  // setup word map
  words := make(map[string]int)
  var found int = 0
  for _,query := range queries {
    for _, word := range w.tokenize(query) {
      words[word] = w.max
      found++
    }
  }

  // walk data
  err := WalkTsvGzCells(fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang), func(row, _ int, cell string) error {
    for word := range words {
      if cell == strings.ToLower(word) {
        words[word] = row
        found--
      }
    }

    if found == 0 { return io.EOF }
    return nil
  })
  if err != nil { return nil, err }

  // build results
  results := make(map[string]int, len(queries))
  for _,query := range queries {
    q := w.tokenize(query)
    fqs := make([]int, len(q))

    for i, word := range q {
      fqs[i] = words[word]
    }

    results[query] = HalfHarmonicMean(fqs)
  }

  return results, nil
}



