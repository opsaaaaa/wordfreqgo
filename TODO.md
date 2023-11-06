

# TODO

# Features
## query
Query a word frequencies from existing dataset

- [x] Lookup, query word frequency function.
- [x] TopN, list top 10-1000 words without frequencies.
- [] Range by Frequency, get a word list with frequency range, like 8.0, 7.0
- [] Range by Rank, return a word list with a rank range. like #10, #20
- [] add option to add custom dataset.
- [] add number handling
- [] add tokenize function for handling Japanese and Chinese multi word queries.
- [] export data function. Iterate over every word with frequency and word.

## Count
Count and build your own word frequency dataset.
- [] count words and build your own datasets of word frequencies.

## Convert
Convert from one frequency unit to another.


# Tasks

- [] separate query and convert into their own packages.
- [] default to cb.
- [] make count package
- [] fix spelling error for assertSame and assertErrNil
- [] make an actual readme with some explanation
- [] write some example usage
- [] explain unit conversion.
- [] explain frequency estimation, bias


## other tasks
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
- [x] consider server operations with a database. and how that should work.
    - it should use a database.
- [] investigate pulishing options.
- [] consider spliting the large/small files. it tries the small file first, and does the other half if it doesnt find the word.
    - that optimization currenlt doesnt matter tho.
- [] calculate frequency for a given input text and tokenizer



# Archive

## Tasks
- [x] convert all the msgpack.gz files into tsv.gz files.
- [x] walk tsv file
- [x] write find the word functions.
- [x] create a git repo
- [x] convert from cb to zipf and other formats
- [x] check what part of the loop im running toLower case on. make sure its not repeated unneesisarilly.
- [x] fix the bug with duplicate queries
- [x] make it go faster
- [x] add top n function
- [x] fix combinations. half harmonic isn't correct. the duke, is shown as more common than duke, cause the is high.
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
- [x] remove wordtree code, as it is slower.

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
- [x] benchmark tree lookup
    - slow


