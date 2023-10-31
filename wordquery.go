package main

import (
	"fmt"
	"strings"
)

// data/small_en.tsv.gz
const DATA_FILE_TSV_GZ = "data/%s_%s.tsv.gz"

type WordQuery struct {
  // en, ja, es, fr
  lang   string
  
  // the size of dataset used.
  // small has upto CB 600 and large has up to CB 800
  // large, small
  size   string

  // the default CB value for words not found in the dataset.
  // between 600~800
  max    int
  
  // tokenize queries for multi-word phrases.
  // its only important if the language doesn't use white space as word boundaries
  // and your querying a multi-word phrase.
  // default is strings.Fields 
  tokenize func(s string) []string

  // describes the bias between the minimum and maximum possible values
  // when estimating multi-word phrases, like "new york".
  // 1 assumes the words are unrelated and 0 assumes the word combo only occurs together
  // 1.0~0.0, 0.5
  comboBias float64
}

func NewWordQuery(lang string) *WordQuery {
  return &WordQuery{
    lang:     lang,
    size:     "small",
    max:      800,
    tokenize: strings.Fields,
    comboBias: 0.5,
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
    minfq := 0

    for i, word := range q {
      fqs[i] = results[word]
      if fqs[i] > minfq { minfq = fqs[i] }
    }

    output[query] = int(float64(minfq) + ((CbAndProbabilities(fqs...) - float64(minfq))*w.comboBias))
  }

  return output, nil
}



