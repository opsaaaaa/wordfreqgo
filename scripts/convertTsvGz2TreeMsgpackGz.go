package main

import (
	"bufio"
	"compress/gzip"
	"github.com/vmihailenco/msgpack/v5"
	"fmt"
	"io"
	"os"
	"strings"
)


type WordNode string
type WordTree map[WordNode]WordTree

func NewWordTree() WordTree {
  return make(WordTree, 0)
}
// store a word within the tree.
func (t *WordTree) Insert(word string, val int32) {
  t.insert([]rune(word), val)
}
func (t *WordTree) insert(word []rune, val int32) {
  m := *t
  var n WordNode
  i := 0

  for ; i < len(word) - 1;i++ {
    n = WordNode(word[i])

    if _, ok := m[n]; !ok {
      m[n] = WordTree{}
    }
    m = m[n]
  }
  n = WordNode(string(word[i]) + string(rune(val)))

  if _, ok := m[n]; !ok {
    m[n] = nil
  }
}

func (t *WordTree) WriteTreeMsgpackGz(filename string) (err error) {
  // Create the file
  f, err := os.Create(filename)
  if err != nil { return }
  defer f.Close()

  // Create a gzip writer
  gw := gzip.NewWriter(f)
  defer gw.Close()

  // Pack the data with msgpack
  data, err := t.Pack()
  if err != nil { return }

  // Write the data to the file
  _, err = gw.Write(data)
  if err != nil { return }

  return nil
}
func (t *WordTree) Pack() ([]byte, error) {
  return msgpack.Marshal(t)
}

func tsvGz2TreeMsgpackGz(src, dest string) {
  wt := NewWordTree()
  err := WalkTsvGzCells(src, func(row, col int, cell string) error {
    if cell != "" {
      wt.Insert(cell, int32(row))
    }
    return nil
  })
  if err != nil { panic(err) }
  wt.WriteTreeMsgpackGz(dest)
  fmt.Println("Complete:", src, dest)
}

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

// func scanBufioGz(filename string, callback func(*bufio.Scanner) error) error {
//   return readGzipGz(filename, func(r *gzip.Reader) error {
//     return callback(bufio.NewScanner(r))
//   })
// }

func WalkTsvGzCells(filename string, callback func(row int, col int, cell string) error) (err error) {
  err = readBufioGz(filename, func(r *bufio.Reader) error {
    row := 0
    for {
      line, err := r.ReadString('\n')
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
  })
  return err
}

func main() {
  tsvGz2TreeMsgpackGz("data/large_ar.tsv.gz", "data/ar.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_bg.tsv.gz", "data/bg.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_bn.tsv.gz", "data/bn.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_ca.tsv.gz", "data/ca.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_cs.tsv.gz", "data/cs.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_da.tsv.gz", "data/da.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_de.tsv.gz", "data/de.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_el.tsv.gz", "data/el.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_en.tsv.gz", "data/en.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_es.tsv.gz", "data/es.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_fa.tsv.gz", "data/fa.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_fi.tsv.gz", "data/fi.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_fil.tsv.gz", "data/fil.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_fr.tsv.gz", "data/fr.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_he.tsv.gz", "data/he.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_hi.tsv.gz", "data/hi.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_hu.tsv.gz", "data/hu.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_id.tsv.gz", "data/id.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_is.tsv.gz", "data/is.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_it.tsv.gz", "data/it.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_ja.tsv.gz", "data/ja.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_ko.tsv.gz", "data/ko.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_lt.tsv.gz", "data/lt.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_lv.tsv.gz", "data/lv.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_mk.tsv.gz", "data/mk.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_ms.tsv.gz", "data/ms.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_nb.tsv.gz", "data/nb.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_nl.tsv.gz", "data/nl.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_pl.tsv.gz", "data/pl.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_pt.tsv.gz", "data/pt.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_ro.tsv.gz", "data/ro.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_ru.tsv.gz", "data/ru.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_sh.tsv.gz", "data/sh.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_sk.tsv.gz", "data/sk.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_sl.tsv.gz", "data/sl.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_sv.tsv.gz", "data/sv.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_ta.tsv.gz", "data/ta.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_tr.tsv.gz", "data/tr.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_uk.tsv.gz", "data/uk.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_ur.tsv.gz", "data/ur.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/small_vi.tsv.gz", "data/vi.tree.msgpack.gz")
  tsvGz2TreeMsgpackGz("data/large_zh.tsv.gz", "data/zh.tree.msgpack.gz")
  fmt.Println("DONE")
}


/*
large_ar.tsv.gz, ar.tree.msgpack.gz, ar, 800
small_bg.tsv.gz, bg.tree.msgpack.gz, bg, 600
large_bn.tsv.gz, bn.tree.msgpack.gz, bn, 800
large_ca.tsv.gz, ca.tree.msgpack.gz, ca, 800
large_cs.tsv.gz, cs.tree.msgpack.gz, cs, 800
small_da.tsv.gz, da.tree.msgpack.gz, da, 600
large_de.tsv.gz, de.tree.msgpack.gz, de, 800
small_el.tsv.gz, el.tree.msgpack.gz, el, 600
large_en.tsv.gz, en.tree.msgpack.gz, en, 800
large_es.tsv.gz, es.tree.msgpack.gz, es, 800
small_fa.tsv.gz, fa.tree.msgpack.gz, fa, 600
large_fi.tsv.gz, fi.tree.msgpack.gz, fi, 800
small_fil.tsv.gz, fil.tree.msgpack.gz, fil, 600
large_fr.tsv.gz, fr.tree.msgpack.gz, fr, 800
large_he.tsv.gz, he.tree.msgpack.gz, he, 800
small_hi.tsv.gz, hi.tree.msgpack.gz, hi, 600
small_hu.tsv.gz, hu.tree.msgpack.gz, hu, 600
small_id.tsv.gz, id.tree.msgpack.gz, id, 600
small_is.tsv.gz, is.tree.msgpack.gz, is, 600
large_it.tsv.gz, it.tree.msgpack.gz, it, 800
large_ja.tsv.gz, ja.tree.msgpack.gz, ja, 800
small_ko.tsv.gz, ko.tree.msgpack.gz, ko, 600
small_lt.tsv.gz, lt.tree.msgpack.gz, lt, 600
small_lv.tsv.gz, lv.tree.msgpack.gz, lv, 600
large_mk.tsv.gz, mk.tree.msgpack.gz, mk, 800
small_ms.tsv.gz, ms.tree.msgpack.gz, ms, 600
large_nb.tsv.gz, nb.tree.msgpack.gz, nb, 800
large_nl.tsv.gz, nl.tree.msgpack.gz, nl, 800
large_pl.tsv.gz, pl.tree.msgpack.gz, pl, 800
large_pt.tsv.gz, pt.tree.msgpack.gz, pt, 800
small_ro.tsv.gz, ro.tree.msgpack.gz, ro, 600
large_ru.tsv.gz, ru.tree.msgpack.gz, ru, 800
small_sh.tsv.gz, sh.tree.msgpack.gz, sh, 600
small_sk.tsv.gz, sk.tree.msgpack.gz, sk, 600
small_sl.tsv.gz, sl.tree.msgpack.gz, sl, 600
large_sv.tsv.gz, sv.tree.msgpack.gz, sv, 800
small_ta.tsv.gz, ta.tree.msgpack.gz, ta, 600
small_tr.tsv.gz, tr.tree.msgpack.gz, tr, 600
large_uk.tsv.gz, uk.tree.msgpack.gz, uk, 800
small_ur.tsv.gz, ur.tree.msgpack.gz, ur, 600
small_vi.tsv.gz, vi.tree.msgpack.gz, vi, 600
large_zh.tsv.gz, zh.tree.msgpack.gz, zh, 800
*/
