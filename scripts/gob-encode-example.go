package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type RuneTree map[rune]interface{}

func WriteMapToFile(filename string, data RuneTree) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadMapFromFile(filename string) (RuneTree, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var data RuneTree
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	// Example usage
	data := RuneTree{
		'a': RuneTree{
			'c': RuneTree{
				'd': 2,
			},
		},
	}

	filename := "nested_map.gob"

	// Write the map to file
	err := WriteMapToFile(filename, data)
	if err != nil {
		fmt.Println("Error writing map to file:", err)
		return
	}
	fmt.Println("Map written to file successfully.")

	// Read the map from file
	loadedData, err := ReadMapFromFile(filename)
	if err != nil {
		fmt.Println("Error reading map from file:", err)
		return
	}
	fmt.Println("Map loaded from file:", loadedData)
}

// package main

// import (
//   "bytes"
//   "encoding/gob"
//   "fmt"
//   "log"
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

//   // Marshal the RuneTree to a buffer.
//   var buf bytes.Buffer
//   enc := gob.NewEncoder(&buf)
//   if err := enc.Encode(tree); err != nil {
//     log.Fatal("Encode error:", err)
//   }

//   // Unmarshal the data from the buffer back into a RuneTree.
//   dec := gob.NewDecoder(&buf)
//   var unmarshaledTree RuneTree
//   if err := dec.Decode(&unmarshaledTree); err != nil {
//     log.Fatal("Decode error:", err)
//   }

//   // Print the unmarshaled RuneTree.
//   fmt.Printf("%#v\n", unmarshaledTree)
// }
