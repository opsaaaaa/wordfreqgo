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

func readGzipGz(filename string, callback func(*gzip.Reader) error) (err error) {
  file, err := os.Open(filename)
  if err != nil { return err }

  defer file.Close()

  gzReader, err := gzip.NewReader(file)
  if err != nil { return err }

  defer gzReader.Close()

  err = callback(gzReader)
  return err
 
}

func readBufioGz(filename string, callback func(*bufio.Reader) error) error {
  return readGzipGz(filename, func(r *gzip.Reader) error {
    return callback(bufio.NewReader(r))
  })
}

func scanBufioGz(filename string, callback func(*bufio.Scanner) error) error {
  return readGzipGz(filename, func(r *gzip.Reader) error {
    return callback(bufio.NewScanner(r))
  })
}

func splitOnTabNewlineSpace(data []byte, atEOF bool) (int,[]byte,error) {
  start := 0
  i := 0
  // skip leading empty lines
  for ; i < len(data); i++ {
    if data[i] != '\t' && data[i] != '\n' && data[i] != ' ' {
      start = i
      break
    }
  }
  // skip leading empty lines
  for ; i < len(data); i++ {
    if data[i] == '\t' || data[i] == '\n' || data[i] == ' ' {
      return i + 1, data[start:i], nil
    }
  }
  if !atEOF { return 0, nil, nil }
  return 0, data, bufio.ErrFinalToken
}


func CellListRangeTsvGz(filename string, start, end int) ([]string, error) {
  out := make([]string,0)

  err := scanBufioGz(filename, func(sc *bufio.Scanner) error {
    sc.Split(splitOnTabNewlineSpace)

    c := 0

    for ; c < start; c++ {
      sc.Scan()
    }
    for sc.Scan() {
      if sc.Err() != nil { return sc.Err() }
      out = append(out, sc.Text())
      c++
      if c >= end {break}
    }
    return sc.Err()
  })

  return out, err
}

// gat all the words between two line numbers
// func ListRangeRowsTsvGz(filename string, start, end int) ([]string, errorr) {
//   err := readBufioGz(filename, func(r *bufio.Reader) error {
//     row := 0

//   })
// }

// find the linenumber/row for a list of words.
func SearchTsvGzRows(filename string, queries []string, defualt int) (map[string]int,error) {
  // setup
  results := make(map[string]int, len(queries))
  for _, q := range queries {
    results[strings.ToLower(q)] = defualt
  }

  notfound := make([]string, len(results))
  for k := range results {
    notfound = append(notfound, k)
  }

  // walk
  err := readBufioGz(filename, func(r *bufio.Reader) error {
    row := 0
    for {
      if len(notfound) <= 0 { break }

      line, err := r.ReadString('\n')

      if err == io.EOF {
        break
      } else if err != nil {
        return err
      } else if line == "" {
        continue
      }

      fields := strings.Split(strings.TrimSpace(line), "\t")

      for _, cell := range fields {
        for i, query := range notfound {
          if cell == query {
            results[query] = row
            notfound = append(notfound[:i], notfound[i+1:]...)
          }
        }
      }
      row++
    }
    return nil
  })


  output := make(map[string]int, len(queries))
  for _, q := range queries {
    output[q] = results[strings.ToLower(q)]
  }

  return output, err
}


// I am removing this for now, because I think it will be better to use ReadTsvGz and implment more speciallized functions.

// func WalkTsvGzCells(filename string, callback func(row int, col int, cell string) error) (err error) {
//   err = ReadTsvGz(filename, func(r *bufio.Reader) error {
//     row := 0
//     for {
//       line, err := r.ReadString('\n')
//       if err == io.EOF {
//         break
//       } else if err != nil {
//         return err
//       }

//       fields := strings.Split(strings.TrimSpace(line), "\t")
//       for col, cell := range fields {
//         err = callback(row, col, cell)
//         if err == io.EOF {
//           break
//         } else if err != nil {
//           return err
//         }
//       }
//       row++
//     }
//     return nil
//   })
//   return err
// }


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

