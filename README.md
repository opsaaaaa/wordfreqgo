

# wordfreqgo, Word Frequency Library for Golang

A library for handling word/token frequencies.

Based on the python library [rspeer/wordfreq](https://github.com/rspeer/wordfreq).

## Notable Alterations
- Ported to golang instead of python.
- Used tsv format for data storage instead of msgpack.
- Estimating the frequencies of multi word phrases functions differently.
- Languages like Chinese I haven't put effort into supporting well.

## Features
- Lookup word Frequencies in various languages.
- TopN word list.
- Convert cb zipf to fpwm and other units.

## Planned Features
- tools for counting word/token frequency and building custom datasets.
- get a word list with frequency range, like 8.0, 7.0
- return a word list with a rank range. like #10, #20
- CLI interface
- number handling



## Word Frequency Units Explained

### Overview

| name        | range           | desciption                                  |
|-------------|-----------------|---------------------------------------------|
| fq          | 1 ~ 0           | frequency represented as a proportion between 0 and 1. Occurrence count divided by total words/tokens in courpous. |
| fpmw        | 1 million to 0  | The number of times the word occurs in one million words. |
| fpbw        | 1 billion to 0  | Number of times it occurs one billion words |
| word rank   | 1+n             | Frequency rank relative to all the other words within your corpus.
| zipf scale  | 9.0 to 0.0-     | Its log10 of frequency per billion words. Named after the American linguist George Kingsley Zipf |
| cb          | 0 to 900+       | Its a word frequency from of logarithmic centibel scale. Basically zipf optimized for computer storage. |


| name        | Advantages      | Disadvantages                     |
|-------------|-----------------|-----------------------------------|
| fq          | simple          | It uses lots and lots of decimals |
| fpmw        | Its straight forward to calculated and understand | Its not easy for humans to compare. for some words its less than 1 |
| fpbw        | Words arn't going to be less than one. | nobody uses it |
| zipf scale  | Easy for humans to compare. | requires decimals for accuracy; It can technically cross 0. |
| cb          | we can safely represent it as a positive integer without sacrificing significant accuracy | Its not widely used |

### Reference Table

The first row represent a scenario were all or 100% of the words in the corpus are the same as the query word.
The last row represents a scenario were the original corpus doesn't contain the query word. Thus it never occurred.
En words column are examples sampled from the dataset at or near each value.


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
|              | -∞     | ∞        | 0%           | 0             | 0         | 0           |


### Where does cb come from?

cb is the word frequency unit used by the dataset pulled from the wordfreq program.
[rspeer/wordfreq](https://github.com/rspeer/wordfreq)
Its very similar to zipf, but with a different scale and 0 point.
For the purposes of this package cb is represented as a positive integer whenever possible.


## Estimating Frequencies for Multi Word Phrases

The dataset only contains single word tokens.
For example "new york" is recorded in the data as two separate tokens "new" and "york".

To handle multi word phrase the package estimates the value from the combined words.
It is an estimate because its unknown how correlated the component words of a phrase are.
the comboBias can be set to bias between (0.0) assuming the component words only occur in the provided phrase,
to (1.0) assuming all the words are completely unrelated.

A phrase can not possibly occur more than its individual words.
Therefor the highest frequent possible for a phrase is the lowest frequency of its component words.
The phrase "new york" can not be more common than the word "york" which it contains.
The estimate uses the lowest word frequency in the phrase as a maximum frequency.

For a minimum frequency we combined the probability of the component words occurring individually using probability math.
`P(A and B and C) ≈ P(A) * P(B) * P(C)` where P(n) is a probability of (n) occurring represented as a proportion between 1 and 0.
This value assumes the words occur like dice rolls which are independent and unrelated.

The estimate takes a value somewhere between the minimum and maximum possible frequencies.
The default bias is 0.5. Which picks a point halfway between the two.
This will over estimate uncorrelated phrases and under estimate correlated phrases.

I'm inclined to think that most phrases provided by a user are likely to be more correlated than not.
But the default bias for now will be 0.5.

I used conjugations to vaguely gauge if estimates are in the right ballpark.
The deviation between the phrases I tested is 0.54.
The actual difference between conjugation frequencies will very significantly.
Note it isn't very accurate and it's not a one to one comparison.
But based on that gauge I am guessing that the estimates are probably close enough.

| a     | b     | c     | a         | b         | c         | a - b  |
|-------|-------|-------|-----------|-----------|-----------|--------|
| 6.33  | 5.985 | 6.14  | it's      | it is     | its       |  0.345 |
| 6.20  | 5.195 | 4.74  | don't     | do not    | dont      |  1.005 |
| 6.15  | 4.825 | 4.69  | I'm       | I am      | Im        |  1.325 |
| 5.80  | 5.305 | 4.35  | can't     | can not   | cant      |  0.495 |
| 5.72  | 5.755 |       | I've      | I have    |           | -0.035 |
| 5.68  | 4.805 | 4.18  | didn't    | did not   | didnt     |  0.875 |
| 5.41  | 5.720 | 3.89  | isn't     | is not    | isnt      | -0.310 |
| 5.36  | 5.315 |       | I'd       | I would   |           |  0.045 |
| 5.28  | 5.600 | 3.69  | wasn't    | was not   | wasnt     | -0.320 |
| 5.18  | 5.115 | 3.49  | wouldn't  | would not | wouldnt   |  0.065 |
| 5.05  | 5.560 |       | aren't    | are not   |           | -0.510 |
| 5.04  | 5.395 |       | we've     | we have   |           | -0.355 |
| 4.57  | 5.115 |       | he'd      | he would  |           | -0.545 |
| 4.44  | 5.040 |       | we'd      | we would  |           | -0.600 |
| 4.19  | 5.245 |       | it'd      | it would  |           | -1.055 |
| 4.18  | 4.895 |       | she'd     | she would |           | -0.715 |


## License


This is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

Any data under `lib/freql/data` comes from the wordfreq repsoitory and is therefore subject to its License

[rspeer/wordfreq#license](https://github.com/rspeer/wordfreq#license)

## Credits


- I read the code from the python wordfreq program, and I'm using the data from the source code. https://github.com/rspeer/wordfq














