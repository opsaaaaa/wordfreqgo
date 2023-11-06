package main

import (
	"fmt"
	"strings"
)

// data/small_en.tsv.gz
const DATA_FILE_TSV_GZ = "data/%s_%s.tsv.gz"
const DATA_FILE_TREE_MSG_GZ = "data/%s.tree.msgpack.gz"


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
  // 1.0 assumes the words are unrelated and 0.0 assumes the word combo only occurs together
  // 0.5 is the default
  comboBias float64

  // the unit
  unitConversion func(n float64) float64
}

func NewWordQuery(lang string) *WordQuery {
  return &WordQuery{
    lang:     lang,
    size:     "small",
    max:      800,
    tokenize: strings.Fields,
    comboBias: 0.5,
    unitConversion: CbToZipf,
  }
}

// func (w *WordQuery) Lookup(query string) (float64, error) {
//   words := w.tokenize(query)

//   results, err := SearchTsvGzRows(fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang), words, w.max)

//   fqs := make([]int, len(words))
//   minfq := 0

//   for i, word := range words {
//     fqs[i] = results[word]
//     if fqs[i] > minfq { minfq = fqs[i] }
//   }

//   return w.calcQueryValue(minfq, fqs...), err
// }

// TODO: add a thing for unit maybe
type WordValFloat struct {
  word string
  val float64
}

// func (w *WordQuery) Lookup(queries ...string) ([]WordValFloat, error) {
//   // for _,query := range queries {
//   //   words = append(words, w.tokenize(query)...)
//   // }

//   wt := NewWordTree()
//   err := wt.ReadTreeMsgpackGz(w.filenameTreeMsgGz())
//   if err != nil { return nil, err }

//   resutls := make([]WordValFloat, len(queries))

//   var fqs []int
//   var words []string
//   minfq := 0

//   for i, query := range queries {

//     words = w.tokenize(strings.ToLower(query))
//     fqs = make([]int, len(words))

//     for j, word := range words {
//       fqs[j] = int(wt.Lookup(word))
//       if fqs[j] > minfq { minfq = fqs[j] }
//       // fqs = append(fqs, int(wt.Lookup(word)))
//     }

//     resutls[i].word = query
//     resutls[i].val = w.calcQueryValue(minfq, fqs...)
//   }

//   return resutls, nil

// }

func (w *WordQuery) Lookup(queries ...string) ([]WordValFloat, error) {

  words := make([]string, 0)
  for _,query := range queries {
    words = append(words, w.tokenize(query)...)
  }

  results, err := SearchTsvGzRows(w.filenameTsvGz(), words)
  if err != nil { return nil, err }

  wordMap := make(map[string]int, len(results))
  for _,wci := range results {
    wordMap[wci.word] = wci.cb
  }

  output := make([]WordValFloat, len(queries))
  for i,query := range queries {
    q := w.tokenize(query)
    fqs := make([]int, len(q))
    minfq := 0

    for i, word := range q {
      fqs[i] = wordMap[word]
      if fqs[i] > minfq { minfq = fqs[i] }
    }

    output[i].val = w.calcQueryValue(minfq, fqs...)
    output[i].word = query
  }

  return output, nil
}

func (w *WordQuery) filenameTreeMsgGz() string {
  return fmt.Sprintf(DATA_FILE_TREE_MSG_GZ, w.lang)
}

func (w *WordQuery) filenameTsvGz() string {
  return fmt.Sprintf(DATA_FILE_TSV_GZ, w.size, w.lang)
}

func (w *WordQuery) calcQueryValue(minfq int, fqs ...int) float64 {
  return w.unitConversion(float64(minfq) + ((CbAndProbabilities(fqs...) - float64(minfq))*w.comboBias))
}

func (w *WordQuery) TopN(n int) ([]string, error) {
  return CellListRangeTsvGz(w.filenameTsvGz(), 0, n)
}


