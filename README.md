
# TODO

- [x] convert all the msgpack.gz files into tsv.gz files.
- [x] walk tsv file
- [x] write find the word functions.
- [x] create a git repo
- [] convert from cb to zipf and other formats
- [] add other functions like top n
- [] extract tsv files function.
- [] extract tsv word list function.
- [] extract work list with frequency ranges.
    - just use line numbers to get the right bits, if from zipf, round min and max with celi and floor.
- [] make it a commandline interface.
- [] consider filtering with spell checker. (maybe another project)
- [] consider server operations with a database. and how that should work.
- [] make an actuall readme with some explination
- [] investigate pulishing options.
- [] consider spliting the large/small files. it tries the small file first, and does the other half if it doesnt find the word.
    - that optimization currenlt doesnt matter tho.
- [x] check what part of the loop im running toLower case on. make sure its not repeated unneesisarilly.
- [x] make it go faster
- [x] fix the bug with duplicate queries
- [] preverve order of query
- [] calculate frequency for a given input text and tokenizer
- [] fix combonations. half harmonic isn't correct. the duke, is shown as more common than duke, cause the is high.
- [] for my usecase I will want to filter by correcly spelled words. common misspellings are included in the dataset.
- [] handle numbers.

- [] filter out things like the dot in mr. dr.
- [x] do correct math.
