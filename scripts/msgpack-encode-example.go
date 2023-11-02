package main

import (
  "fmt"
  "github.com/vmihailenco/msgpack/v4"
)

type RuneTree map[rune]interface{}

func main() {
  // Example deeply nested map
  nestedMap := RuneTree{
    'a': RuneTree{
      'b': RuneTree{
        'c': 123,
      },
      'c': 2,
    },
  }
  // it doens't like mixed types on the same layer.
  // to use msg pack I'd have to not mix layers without telling it.
  //  {a: {children: {...}, val: rune}, ...}
  // So I'd have to do a ligit tree anyway.
  // the whole point of this is to store indexes compactly.
  // idk if that is going to work.
  // Best to use a database I think.
  // [rune, []]
  // [a: [ b: [c: 1] ...] ...]
  // RuneBranch
  // RuneLeaf

  // Marshal the nested map into msgpack format
  msgpackData, err := msgpack.Marshal(nestedMap)
  if err != nil {
    fmt.Println("Error marshaling:", err)
    return
  }

  fmt.Println("Marshalled data:", msgpackData)

  // Unmarshal the msgpack data back into a map
  var unmarshaledMap RuneTree
  err = msgpack.Unmarshal(msgpackData, &unmarshaledMap)
  if err != nil {
    fmt.Println("Error unmarshaling:", err)
    return
  }

  fmt.Println("Unmarshaled map:", unmarshaledMap)
}

// package main

// import (
//   "fmt"
//   "log"

//   "github.com/vmihailenco/msgpack/v5"
// )

// // RuneTree represents a deeply nested map with runes as keys and either RuneTree or int32 as values.
// type RuneTree map[rune]interface{}

// func main() {
//   // Create a sample RuneTree.
//   tree := RuneTree{
//     'a': RuneTree{
//       'b': int32(42),
//       'c': RuneTree{
//         'd': int32(99),
//       },
//     },
//   }

//   // Marshal the RuneTree to MessagePack.
//   marshaledData, err := msgpack.Marshal(tree)
//   if err != nil {
//     log.Fatal("Marshal error:", err)
//   }

//   // Unmarshal the MessagePack data back into a RuneTree.
//   var unmarshaledTree RuneTree
//   err = msgpack.Unmarshal(marshaledData, &unmarshaledTree)
//   if err != nil {
//     log.Fatal("Unmarshal error:", err)
//   }

//   // Print the unmarshaled RuneTree.
//   fmt.Printf("%#v\n", unmarshaledTree)
// }
