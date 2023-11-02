
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

## indexes
This is a very fun tangient but this is infact a tangent.
it is good enough for what it is. I should honestly use a database for fast lookup.

- [x] tinker with encoding/gob
    - see if it will work for saving indexes
    - https://pkg.go.dev/encoding/gob
- [x] thinker with msgpack
    - to see if it will work for saving indexes.
- [] make it faster with indexes. use a tree index for compactness
    - map[rune]map[rune]...int
    - https://stackoverflow.com/questions/44971026/get-value-form-nested-map-type-mapstringinterface
    - key, value
    - [[rune, val]]
    - jk we're noping this feature.

## other tasks
- [] get range by line
- [] preverve order of query
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
- [] make it a commandline interface.
- [] make an actuall readme with some explination

## consider decisions
- [] consider filtering with spell checker. (maybe another project)
- [] consider server operations with a database. and how that should work.
- [] investigate pulishing options.
- [] consider spliting the large/small files. it tries the small file first, and does the other half if it doesnt find the word.
    - that optimization currenlt doesnt matter tho.
- [] calculate frequency for a given input text and tokenizer



# thinking

we could go the pointer route.

type RuneTree struct {
    val rune
    children []runeTree
}

and children would be sorted so that we can use cools search functions.


