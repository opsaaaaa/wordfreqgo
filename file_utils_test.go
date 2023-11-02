package main

import (
  "testing"
)

func Test_SearchTsvGz(t *testing.T) {
  // line numbers start at 0
  expected := map[string]int{"a2":1,"b2":1,"a6":5,"bad":9}
  results, err := SearchTsvGzRows("test/sample_test.tsv.gz", []string{"a2", "b2", "a6", "bad"})
  assetErrNil(t, err)
  for _,wci := range results {
    asertSame(t, wci.cb, expected[wci.word])
  }
}

func Test_CellListRangeTsvGz(t *testing.T) {
  expected := []string{"b2","c2","a3"}
  results, err := CellListRangeTsvGz("test/sample_test.tsv.gz", 1, 4)
  assetErrNil(t, err)
  if len(results) != 3 { 
    t.Fatalf("Expected len to equal 3 and got %v:\n%v\n",len(results), results)
  }
  for i,v := range expected {
    asertSame(t, v, results[i])
  }
}

func Test_CellListRangeTsvGz_overflow(t *testing.T) {
  expected := []string{"a2","b2","c2","a3"}
  results, err := CellListRangeTsvGz("test/sample_test.tsv.gz", 0, 100)
  assetErrNil(t, err)
  if len(results) != 9 { 
    t.Fatalf("Expected len to equal 3 and got %v:\n%v\n",len(results), results)
  }
  for i,v := range expected {
    asertSame(t, v, results[i])
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


// a2 b2  c2
// a3
// a4

// a6 b6


// a9
