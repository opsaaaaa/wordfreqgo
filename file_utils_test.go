
package main

import (
  "testing"
)

func Test_WalkTsvGzCells(t *testing.T) {
  expectedCells := [20][8]string{
    {},
    {"a2", "b2", "c2"},
    {"a3"},
    {"a4"},
    {},
    {"a6","b6"},
    {},
    {},
    {"a9"},
    {},
  }

  err := WalkTsvGzCells("test/sample_test.tsv.gz", func(row int, col int, cell string) error {
    if expectedCells[row][col] != cell {
      t.Errorf("Expected '%s', but got '%s'", expectedCells[row][col], cell)
    }
    return nil
  })
  if err != nil {
    t.Errorf("Error walking TSV file: %v", err)
  }
}


// a2	b2	c2
// a3
// a4

// a6	b6


// a9
