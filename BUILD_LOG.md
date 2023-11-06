

# The current status.

English and similar languages with white space word boundaries should function fine.
The wordquery struct accepts a tokenize function.
So other types of boundaries can be supported that way.

I suspect Chinese and Japanese will need more work to function correctly.
Other languages may also need some attention.
I am making that a low priority as it is complex and not applicable to my use case.

Numbers don't yet work correctly.
Again low priority for me.

I would like to have a function to easily export data.
Specifically to write it to a database.
I think I'll write that in the project I use that feature in.
Then I'll decide if I am going to add that tooling here or not.

The next things to be DONE.
- write a more useful readme
- consider how to publish the package. I'll do that when I need to use the package in another project.
- make a command line utility. Perhaps a wrapper package called "wordfreqgo-cli" ?.


# Removal of word tree code.

I removed the word tree code and several experiments that aren't needed in the final product.
That code is preserved in a branch called "2023-11-06".

This code includes scripts to convert the source msgpack word data to the tsv format. (from the wordfreq py libary.)



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


