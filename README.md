
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
- [] convert from cb to zipf and other formats
- [] write a baller test for unit conversions.
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
