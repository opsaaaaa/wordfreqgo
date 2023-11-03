package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

func readFreqBinData(data_file string) ([][]string, error) {
  // Open the gzipped file
  data, err := os.ReadFile(data_file)
  if err != nil {
    return nil, err
  }

  // Unzip the gzipped file
  gr, err := gzip.NewReader(bytes.NewBuffer(data))
  if err != nil {
    return nil, err
  }
  defer gr.Close()

  // Read the unzipped data
  unzippedData, err := io.ReadAll(gr)
  if err != nil {
    return nil, err
  }

  // Unmarshal the msgpack data into a slice of interface{}
  var result []interface{}
  err = msgpack.Unmarshal(unzippedData, &result)
  if err != nil {
    return nil, err
  }

  // Convert the rest of the items to []string
  var body [][]string
  for _, item := range result[1:] {
    switch v := item.(type) {
    case []interface{}:
      var line []string
      for _, element := range v {
        line = append(line, element.(string))
      }
      body = append(body, line)
    default:
      return nil, fmt.Errorf("expected a slice for an item, got %T", v)
    }
  }

  return body, nil
}

func writeToTSV(data [][]string, filename string) error {
  // Create the file
  f, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer f.Close()

  // Create a gzip writer
  gw := gzip.NewWriter(f)
  defer gw.Close()

  // Write the data to the file
  for _, line := range data {
    _, err := gw.Write([]byte(strings.Join(line, "\t") + "\n"))
    if err != nil {
      return err
    }
  }

  return nil
}


func convertMsgpack2Tsv(msgpackFile string, tsvFile string) {
  data, err := readFreqBinData(msgpackFile)
  if err != nil { log.Fatal(err)}
  err = writeToTSV(data, tsvFile)
  if err != nil { log.Fatal(err)}
}

func main() {
  convertMsgpack2Tsv("data/large_ar.msgpack.gz", "data/large_ar.tsv.gz")
  convertMsgpack2Tsv("data/large_bn.msgpack.gz", "data/large_bn.tsv.gz")
  convertMsgpack2Tsv("data/large_ca.msgpack.gz", "data/large_ca.tsv.gz")
  convertMsgpack2Tsv("data/large_cs.msgpack.gz", "data/large_cs.tsv.gz")
  convertMsgpack2Tsv("data/large_de.msgpack.gz", "data/large_de.tsv.gz")
  convertMsgpack2Tsv("data/large_en.msgpack.gz", "data/large_en.tsv.gz")
  convertMsgpack2Tsv("data/large_es.msgpack.gz", "data/large_es.tsv.gz")
  convertMsgpack2Tsv("data/large_fi.msgpack.gz", "data/large_fi.tsv.gz")
  convertMsgpack2Tsv("data/large_fr.msgpack.gz", "data/large_fr.tsv.gz")
  convertMsgpack2Tsv("data/large_he.msgpack.gz", "data/large_he.tsv.gz")
  convertMsgpack2Tsv("data/large_it.msgpack.gz", "data/large_it.tsv.gz")
  convertMsgpack2Tsv("data/large_ja.msgpack.gz", "data/large_ja.tsv.gz")
  convertMsgpack2Tsv("data/large_mk.msgpack.gz", "data/large_mk.tsv.gz")
  convertMsgpack2Tsv("data/large_nb.msgpack.gz", "data/large_nb.tsv.gz")
  convertMsgpack2Tsv("data/large_nl.msgpack.gz", "data/large_nl.tsv.gz")
  convertMsgpack2Tsv("data/large_pl.msgpack.gz", "data/large_pl.tsv.gz")
  convertMsgpack2Tsv("data/large_pt.msgpack.gz", "data/large_pt.tsv.gz")
  convertMsgpack2Tsv("data/large_ru.msgpack.gz", "data/large_ru.tsv.gz")
  convertMsgpack2Tsv("data/large_sv.msgpack.gz", "data/large_sv.tsv.gz")
  convertMsgpack2Tsv("data/large_uk.msgpack.gz", "data/large_uk.tsv.gz")
  convertMsgpack2Tsv("data/large_zh.msgpack.gz", "data/large_zh.tsv.gz")
  convertMsgpack2Tsv("data/small_ar.msgpack.gz", "data/small_ar.tsv.gz")
  convertMsgpack2Tsv("data/small_bg.msgpack.gz", "data/small_bg.tsv.gz")
  convertMsgpack2Tsv("data/small_bn.msgpack.gz", "data/small_bn.tsv.gz")
  convertMsgpack2Tsv("data/small_ca.msgpack.gz", "data/small_ca.tsv.gz")
  convertMsgpack2Tsv("data/small_cs.msgpack.gz", "data/small_cs.tsv.gz")
  convertMsgpack2Tsv("data/small_da.msgpack.gz", "data/small_da.tsv.gz")
  convertMsgpack2Tsv("data/small_de.msgpack.gz", "data/small_de.tsv.gz")
  convertMsgpack2Tsv("data/small_el.msgpack.gz", "data/small_el.tsv.gz")
  convertMsgpack2Tsv("data/small_en.msgpack.gz", "data/small_en.tsv.gz")
  convertMsgpack2Tsv("data/small_es.msgpack.gz", "data/small_es.tsv.gz")
  convertMsgpack2Tsv("data/small_fa.msgpack.gz", "data/small_fa.tsv.gz")
  convertMsgpack2Tsv("data/small_fil.msgpack.gz", "data/small_fil.tsv.gz")
  convertMsgpack2Tsv("data/small_fi.msgpack.gz", "data/small_fi.tsv.gz")
  convertMsgpack2Tsv("data/small_fr.msgpack.gz", "data/small_fr.tsv.gz")
  convertMsgpack2Tsv("data/small_he.msgpack.gz", "data/small_he.tsv.gz")
  convertMsgpack2Tsv("data/small_hi.msgpack.gz", "data/small_hi.tsv.gz")
  convertMsgpack2Tsv("data/small_hu.msgpack.gz", "data/small_hu.tsv.gz")
  convertMsgpack2Tsv("data/small_id.msgpack.gz", "data/small_id.tsv.gz")
  convertMsgpack2Tsv("data/small_is.msgpack.gz", "data/small_is.tsv.gz")
  convertMsgpack2Tsv("data/small_it.msgpack.gz", "data/small_it.tsv.gz")
  convertMsgpack2Tsv("data/small_ja.msgpack.gz", "data/small_ja.tsv.gz")
  convertMsgpack2Tsv("data/small_ko.msgpack.gz", "data/small_ko.tsv.gz")
  convertMsgpack2Tsv("data/small_lt.msgpack.gz", "data/small_lt.tsv.gz")
  convertMsgpack2Tsv("data/small_lv.msgpack.gz", "data/small_lv.tsv.gz")
  convertMsgpack2Tsv("data/small_mk.msgpack.gz", "data/small_mk.tsv.gz")
  convertMsgpack2Tsv("data/small_ms.msgpack.gz", "data/small_ms.tsv.gz")
  convertMsgpack2Tsv("data/small_nb.msgpack.gz", "data/small_nb.tsv.gz")
  convertMsgpack2Tsv("data/small_nl.msgpack.gz", "data/small_nl.tsv.gz")
  convertMsgpack2Tsv("data/small_pl.msgpack.gz", "data/small_pl.tsv.gz")
  convertMsgpack2Tsv("data/small_pt.msgpack.gz", "data/small_pt.tsv.gz")
  convertMsgpack2Tsv("data/small_ro.msgpack.gz", "data/small_ro.tsv.gz")
  convertMsgpack2Tsv("data/small_ru.msgpack.gz", "data/small_ru.tsv.gz")
  convertMsgpack2Tsv("data/small_sh.msgpack.gz", "data/small_sh.tsv.gz")
  convertMsgpack2Tsv("data/small_sk.msgpack.gz", "data/small_sk.tsv.gz")
  convertMsgpack2Tsv("data/small_sl.msgpack.gz", "data/small_sl.tsv.gz")
  convertMsgpack2Tsv("data/small_sv.msgpack.gz", "data/small_sv.tsv.gz")
  convertMsgpack2Tsv("data/small_ta.msgpack.gz", "data/small_ta.tsv.gz")
  convertMsgpack2Tsv("data/small_tr.msgpack.gz", "data/small_tr.tsv.gz")
  convertMsgpack2Tsv("data/small_uk.msgpack.gz", "data/small_uk.tsv.gz")
  convertMsgpack2Tsv("data/small_ur.msgpack.gz", "data/small_ur.tsv.gz")
  convertMsgpack2Tsv("data/small_vi.msgpack.gz", "data/small_vi.tsv.gz")
  convertMsgpack2Tsv("data/small_zh.msgpack.gz", "data/small_zh.tsv.gz")
}

