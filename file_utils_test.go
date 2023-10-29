package main

import (
	"testing"
)

func Test_SearchTsvGz(t *testing.T) {
  // line numbers start at 0
  expected := map[string]int{"a2":1,"b2":1,"a6":5,"bad":10}
  results, err := SearchTsvGzRows("test/sample_test.tsv.gz", []string{"a2", "b2", "a6", "bad"}, 10)
  assetErrNil(t, err)
  for k,v := range expected {
    asertSame(t, results[k], v)
  }
}

// func Test_WalkTsvGzCells(t *testing.T) {
//   expectedCells := [20][8]string{
//     {},
//     {"a2", "b2", "c2"},
//     {"a3"},
//     {"a4"},
//     {},
//     {"a6","b6"},
//     {},
//     {},
//     {"a9"},
//     {},
//   }

//   err := WalkTsvGzCells("test/sample_test.tsv.gz", func(row int, col int, cell string) error {
//     if expectedCells[row][col] != cell {
//       t.Errorf("Expected '%s', but got '%s'", expectedCells[row][col], cell)
//     }
//     return nil
//   })
//   if err != nil {
//     t.Errorf("Error walking TSV file: %v", err)
//   }
// }


// a2	b2	c2
// a3
// a4

// a6	b6


// a9
