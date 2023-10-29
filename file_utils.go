package main

import (
  "bufio"
  "compress/gzip"
  "io"
  "os"
  "strings"
)

/*
  Way are you not using encoding/csv !?!!
  good question.

  Its because encodeing/csv ignores leading empty lines.
  this is a big problem because I am using line numbers to encode information.
  the other solution to this issue is to add a throw away entery at the top of every the tsv file. (a single tab sould work)
  Right now I am going with a manual parsing solution.

  The downsides is that escaped and quoted characters are not handled correctly.
  This is probably fine because the seperator is a tab and we are only reading known sources within our contoll.

  Why encode information with tsv line numbers?
  Thats a long sotry. tldr wordfreq makers have big brains.
  The important bit here is because I got tired of trying to read msgpacks and having msgpack tools fail to work.
  Technically this format is more space effecient. Anyway. and human readable.
  smh

  I've left some commented out code incase I ever decide to use encoding/csv instead.
*/

func WalkTsvGzCells(filename string, callback func(row int, col int, cell string) error) (err error) {
  file, err := os.Open(filename)
  if err != nil { return err }

  defer file.Close()

  gzReader, err := gzip.NewReader(file)
  if err != nil { return err }

  defer gzReader.Close()

  reader := bufio.NewReader(gzReader)

  row := 0
  for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
      break
    } else if err != nil {
      return err
    }

    fields := strings.Split(strings.TrimSpace(line), "\t")
    for col, cell := range fields {
      err = callback(row, col, cell)
      if err == io.EOF {
        break
      } else if err != nil {
        return err
      }
    }
    row++
  }

  return nil
}


// func WalkTsvGzCells(filename string, callback func(row int, col int, cell string) error) (err error) {
//   f, err := os.Open(filename)
//   if err != nil {
//   return err
//   }
//   defer f.Close()

//   gz, err := gzip.NewReader(f)
//   if err != nil {
//   return err
//   }
//   defer gz.Close()

//   cr := csv.NewReader(bufio.NewReader(gz))
//   cr.Comma = '\t' // Set the delimiter to tab since this is a TSV file.
//   cr.TrimLeadingSpace = false

//   row := 0
//   for {
//   record, err := cr.Read()
//   if err != nil {
//     break
//   }

//   for col, cell := range record {
//     err := callback(row, col, cell)
//     if err != nil {
//     return err
//     }
//   }
//   row++
//   }

//   return nil
// }

