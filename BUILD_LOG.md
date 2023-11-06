

# Benchmark for tree lookup ends up being slower than linear lookup

Using a tree based lookup ends up being slower.
I assume it is becuase of encoding and decoding msgpack data into maps.
or Marshal/Unmarshal.
At some point that has to write every value from the data into the tree.
Where as the tsv linear lookup doens't have to write the whole thing into memory.
Hu, thats very education.

So the trouble with a tree structure, is storing it as a plain file and reading it.
This is where database would win out with lookup speed.
So if lots of queries are needed, then a db sould be the faster choice.
Thats kinda ovbous.

# Thinking About Tree lookup speed.

the issue is the data is stored at the ending key.
So inorder to get that value we need to walk the hash.
But the maximum amount of nodes we have to walk is the number of unique characters in that layer.
Plus any ending nodes in that layer.
That should still be far better than the liner tsv search.
And probably useable on a server.

Technially we should beable to build one form the other.
So having both the word tree and the tsv is redundnat.
The word tree should be smaller.

A topN lookup is way fasfter with the tsv.
Where as with the Tree you have to walk the whole thing.

# Thinking about trees

we could go the pointer route.

type RuneTree struct {
    val rune
    children []runeTree
}

and children would be sorted so that we can use cools search functions.


