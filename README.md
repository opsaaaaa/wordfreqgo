
# Unit Reference Table
♾️


| En Words     | Zipf   | Cb (abs) | Percentage % | Fq Propotion  | Fpmw      | Fpbw        |
|--------------|--------|----------|--------------|---------------|-----------|-------------|
|              | 9.0    | 0        | 100%         | 1             | 1000000   | 1000000000  |
| the          | 8.0    | 100      | 10%          | 0.1           | 100000    | 100000000   |
| that, for    | 7.0    | 200      | 1%           | 0.01          | 10000     | 10000000    |
| said, way    | 6.0    | 300      | 0.1%         | 0.001         | 1000      | 1000000     |
| radio, plans | 5.0    | 400      | 0.01%        | 0.0001        | 100       | 100000      |
| prizes, bail | 3.0    | 500      | 0.001%       | 0.00001       | 10        | 10000       |
| sparing      | 2.0    | 600      | 0.0001%      | 0.000001      | 1         | 1000        |
| cryptology   | 1.0    | 700      | 0.00001%     | 0.0000001     | 0.1       | 100         |
| microcapsule | 0      | 800      | 0.000001%    | 0.00000001    | 0.01      | 10          |
|              | -1     | 900      | 0.0000001%   | 0.000000001   | 0.001     | 1           |
|              | -2     | 1000     | 0.00000001%  | 0.0000000001  | 0.0001    | 0.1         |
|              | -3     | 1100     | 0.000000001% | 0.00000000001 | 0.00001   | 0.01        |
|              | ∞      | ∞        | 0%           | 0             | 0         | 0           |



# TODO

- [x] convert all the msgpack.gz files into tsv.gz files.
- [x] walk tsv file
- [x] write find the word functions.
- [x] create a git repo
- [x] convert from cb to zipf and other formats
- [x] check what part of the loop im running toLower case on. make sure its not repeated unneesisarilly.
- [x] fix the bug with duplicate queries
- [x] make it go faster
- [x] add top n function
- [x] fix combonations. half harmonic isn't correct. the duke, is shown as more common than duke, cause the is high.
- [x] do correct math.
- [x] preserve order of query
- [x] add wordtree code to repo
- [x] Test to make sure wordtree encodes and decodes values correctly
- [x] write a script to convert files to wordtree format.
- [x] write wordquery with tree.
- [x] benchmark wordtree lookup
    - WAY SLOWER, BECAUSE OF WRITTING TREE TO MEMORY FOR LOOKUP.
- [x] check if wordtree file is smaller than tsv
    - hardly smaller, not meaningful
- [] fix spelling error for assertSame and assertErrNil
- [] remove wordtree code, as it is slower.

## indexes
This is a very fun tangient but this is infact a tangent.
it is good enough for what it is. I should honestly use a database for fast lookup.

- [x] tinker with encoding/gob
    - see if it will work for saving indexes
    - https://pkg.go.dev/encoding/gob
- [x] thinker with msgpack
    - to see if it will work for saving indexes.
- [x] make it faster with indexes. use a tree index for compactness
    - map[rune]map[rune]...int
    - https://stackoverflow.com/questions/44971026/get-value-form-nested-map-type-mapstringinterface
    - key, value
    - [[rune, val]]
    - jk we're noping this feature.
    - https://tufin.medium.com/printing-ascii-trees-in-golang-8ae6692fabe0
- [x] check if a tree lookup is smaller than a bin wordlist.
    - just bearly smaller.
- [x] use wordtree code to files.
- [] benchmark tree lookup

## other tasks
- [] make an actuall readme with some explination
    - copy it from my ruby project.
- [] make it a commandline interface.
- [] get range by line
- [] handle numbers.
- [] filter out things like the dot in mr. dr.
    - careful about filtering out puncuation.
    - more than one language here.
    - inject a filter function.
    - only needed on word list queries.
- [] write a baller test for unit conversions.
- [] extract tsv files function.
- [] extract tsv word list function.
- [] extract work list with frequency ranges.
    - just use line numbers to get the right bits, if from zipf, round min and max with celi and floor.

## consider decisions
- [] consider filtering with spell checker. (maybe another project)
- [] consider server operations with a database. and how that should work.
- [] investigate pulishing options.
- [] consider spliting the large/small files. it tries the small file first, and does the other half if it doesnt find the word.
    - that optimization currenlt doesnt matter tho.
- [] calculate frequency for a given input text and tokenizer

# Benchmark for tree lookup end up being slower than linear lookup

Using a tree based lookup ends up being slower.
I assume it is becuase of encoding and decoding msgpack data into maps.
Marshal/Unmarshal.
At some point that has to write every value into the tree.
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


