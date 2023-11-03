package main

import (
  "fmt"
  "github.com/vmihailenco/msgpack/v5"
)




type Node string
type Tree map[Node]Tree
// the other option is to use a string
// and then encode the ending with a length 2 string
// and interpret the final character as the value
// that might be more space efficient.
// cause this way every single entry acts as 2 characters.
// when we only need the ending entries to act as characters.
// Ok I need to get going now. Enough coding.

func main() {
  t := make(Tree, 0)
  t.Insert("one", 1)
  t.Insert("two", 2)
  t.Insert("three", 3)
  t.Insert("four", 4)
  t.Insert("five", 5)
  t.Insert("six", 6)
  t.Insert("seven", 7)
  t.Insert("eight", 8)
  t.Insert("nine", 9)
  t.Insert("ten", 10)
  t.Insert("billy", 100)
  t.Insert("bob", 200)
  t.Insert("frank", 800)


  data, err := msgpack.Marshal(t)
  if err != nil { panic(err) }

  var newtree Tree
  err = msgpack.Unmarshal(data, &newtree)
  if err != nil { panic(err) }

  t.Print()
  newtree.Print()
  fmt.Println(t.Lookup("bob"))
  fmt.Println(newtree.Lookup("bob"))

}

// store a word within the tree.
func (t *Tree) Insert(word string, val int32) {
  t.insert([]rune(word), val)
}
func (t *Tree) insert(word []rune, val int32) {
  m := *t
  var n Node
  i := 0

  for ; i < len(word) - 1;i++ {
    n = Node(word[i])

    if _, ok := m[n]; !ok {
      m[n] = Tree{}
    }
    m = m[n]
  }
  n = Node(string(word[i]) + string(rune(val)))

  if _, ok := m[n]; !ok {
    m[n] = nil
  }
}


// func node(letter rune, i int32) [2]rune {
//   return [2]rune{letter, i}
// }
func (t *Tree) Lookup(word string) int32 {
  return t.lookup([]rune(word))
}
func (t *Tree) lookup(word []rune) int32 {
  m := *t
  i := 0
  for ; i < len(word) - 1;i++ {
    if newMap, ok := m[Node(word[i])]; ok {
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
  // fmt.Println("bar yo")
  return -1
}

func (t *Tree) Print() {
  t.Walk(func(s string, i int) {
    fmt.Printf("%v: %v\n", s, i)
  })
}

func (t *Tree) Walk(callback func(string, int)) {
  t.walkR(make([]rune, 0), callback)
}

func (t *Tree) walkR(acc []rune, callback func(string, int)) {
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
