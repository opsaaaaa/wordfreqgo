package main

import (
	"fmt"
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

  results, err := SearchTsvGzRows(fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang), words, w.max)

  fqs := make([]int, len(words))

  for i, word := range words {
    fqs[i] = results[word]
  }

  return HalfHarmonicMeanArr(fqs), err
}


func (w *WordQuery) LookupMultiple(queries []string) (map[string]int, error) {

  words := make([]string, 0)
  for _,query := range queries {
    words = append(words, w.tokenize(query)...)
  }

  results, err := SearchTsvGzRows(fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang), words, w.max)
  if err != nil { return results, err }

  output := make(map[string]int, len(queries))
  for _,query := range queries {
    q := w.tokenize(query)
    fqs := make([]int, len(q))

    for i, word := range q {
      fqs[i] = results[word]
    }

    output[query] = HalfHarmonicMeanArr(fqs)
  }

  return output, nil
}



