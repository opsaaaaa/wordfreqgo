package main

import "fmt"




type Node [2]rune
type Tree map[Node]Tree
// the other option is to use a string
// and then encode the ending with a length 2 string
// and interpret the final character as the value
// that might be more space efficient.
// cause this way every single entry acts as 2 characters.
// when we only need the ending entries to act as characters.
// Ok I need to get going now. Enough coding.

func main() {
  // t := Tree{
  //   Node{'a',-1}: Tree{
  //     Node{'b',1}: nil,
  //   },
  //   Node{'c',-1}: nil,
  // }
  t := make(Tree, 0)
  t.Insert("bill", 10)
  t.Insert("billy", 11)
  t.Insert("bob", 20)
  t.Insert("frank", 40)
  // fmt.Println(t)
  t.Print()
}

// func node(letter rune, i int32) [2]rune {
//   return [2]rune{letter, i}
// }
func (t *Tree) Lookup(word string) {
  m := *t
  i := 0
  for ; i < len(word) - 1;i++ {
    m = m[Node{r,-1}]
  }
  for i,r := range word {
    
  }
}

// storing words backwards might be more efficient
// The lookup function 
func (t *Tree) LookupR(word string)  {
}

func (t *Tree) Insert(word string, val int32) {
  m := *t
  var n Node
  i := 0
  for ; i < len(word) - 1;i++ {
    n = Node{rune(word[i]), -1}

    if _, ok := m[n]; !ok {
      m[n] = Tree{}
    }
    m = m[n]
  }
  n = Node{rune(word[i]), val}

  if _, ok := m[n]; !ok {
    m[n] = nil
  }
}

func (t *Tree) Print() {
  t.PrintR("")
}

func (t *Tree) PrintR(acc string) {
  for n, m := range *t {
    // acc = acc + string(n[0])
    // not -1 marks a valid word
    if n[1] != -1 {
      fmt.Printf("%v: %v\n",acc + string(n[0]),n[1])
    }
    // 
    if m != nil {
      m.PrintR(acc + string(n[0]))
    }
  }
}

