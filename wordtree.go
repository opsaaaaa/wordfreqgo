package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"github.com/vmihailenco/msgpack/v5"
)

/*
  Word tree is a tree shaped lookup for word frequencies.
  
  cat,3 > {c: {a: {t(3): nil}}}
  cats,9 > {c: {a: {t(3): nil, t: {s(9): nil}}}}

  (0) means the number is encoded as a rune.

  Initially 200 caused an issues because it used more than 2 bytes
  That issue has been fixed by first converting the string to a []rune 

  The test should cover storing values 0-900.
  Just to make sure encoding numbers as runes dones't break for any abatray values.
*/ 


type WordNode string
type WordTree map[WordNode]WordTree

func NewWordTree() WordTree {
  return make(WordTree, 0)
}

// find a value within the tree
func (t *WordTree) Lookup(word string) int32 {
  return t.lookup([]rune(word))
}
func (t *WordTree) lookup(word []rune) int32 {
  m := *t
  i := 0
  for ; i < len(word) - 1;i++ {
    if newMap, ok := m[WordNode(word[i])]; ok {
      m = newMap
    } else {
      return -1
    }
  }

  for n := range m {
    chars := []rune(n)
    if len(chars) == 2 && chars[0] == word[i] {
      return int32(chars[1])
    }
  }
  return -1
}

func (t *WordTree) BuildFromTsv(filename string) error {
  return nil
}

func (t *WordTree) ReadTreeMsgpackGz(filename string) error {
  var data []byte
  err := readGzipGz(filename, func(r *gzip.Reader) (err error) {
    data, err = io.ReadAll(r)
    return err
  })
  if err != nil { return err }
  err = t.Unpack(data)
  return err
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
func (t *WordTree) Unpack(data []byte) error {
  return msgpack.Unmarshal(data, &t)
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

// Walk over all the word values of the tree.
func (t *WordTree) Walk(callback func(string, int)) {
  t.walkR(make([]rune, 0), callback)
}

func (t *WordTree) walkR(acc []rune, callback func(string, int)) {
  for n, m := range *t {
    chars := []rune(n)

    if len(chars) == 2 {
      callback(string(acc)+string(chars[0]), int(chars[1]))
    }

    if m != nil {
      m.walkR(append(acc, chars...), callback)
    }
  }
}

// Print all the word values of the tree.
func (t *WordTree) Print() {
  t.Walk(func(s string, i int) {
    fmt.Printf("%v: %v\n", s, i)
  })
}


