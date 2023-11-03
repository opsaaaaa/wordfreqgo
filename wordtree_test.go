package main

import "testing"

// the only thing this test doesn't cover is t.Print() which is mostly for debugging.
func Test_WordTree_Insert_Lookup_Pack_Unpack(t *testing.T) {
  wq := NewWordQuery("en")
  wq.unitConversion = CbToCb
  top900words, err := wq.TopN(900)
  assetErrNil(t, err)

  wt1 := NewWordTree()
  for i, w := range top900words {
    wt1.Insert(w, int32(i))
  }
  err = wt1.WriteTreeMsgpackGz("/tmp/en.tree.msgpack.gz")
  assetErrNil(t, err)

  wt2 := NewWordTree()
  err = wt2.ReadTreeMsgpackGz("/tmp/en.tree.msgpack.gz")
  assetErrNil(t, err)

  for i, w := range top900words {
    asertSame(t,int32(i), wt2.Lookup(w))
    asertSame(t,int32(i), wt1.Lookup(w))
  }
  asertSame(t,int32(-1), wt2.Lookup("nonexistantentry"))
  asertSame(t,int32(-1), wt1.Lookup("nonexistantentry"))

  wt2.Walk(func(s string, i int) {
    if s == "" || i < 0 || i > 900 {
      t.Fatal("walk failed")
    }
  })
}

