

# wordfreqgo, Word Freqency Libary for Golang

A library for handling word/token frequencies.

Based on the python libary [rspeer/wordfreq](https://github.com/rspeer/wordfreq).


## Features
- Lookup word Frequencies in various languages.
- TopN word list.
- Convert cb zipf to fpwm and other units.
- TODO: tools for counting word/token frequency and building custom datasets.


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

### Refernece Table

The first row represent a senatio were all or 100% of the words in the courpus are the same as the query word.
The last row represents a senatio were the original courpus doens't contain the query word. Thus it never occurd.
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

cb is the word frequency unit used by our initial dataset pulled from the wordfreq program.
[rspeer/wordfreq](https://github.com/rspeer/wordfreq)
Its very similar to zipf, but with a different scale and 0 point.
For the porposes of this package cb is represented as a postive integer whenever possibe.


## Estimating Freqencies for Multi Word Phrases

The dataset only contains single word tokens.
For example "new york" is recorded in the data as two separate tokens "new" and "york".

To handle multi word phrase the package estimates the value from the combined words.
It is an estimate because its unknown how correlated the component words of a phrase are.
the comboBias can be set to bias between (0.0) assuming the component words only occur in the provided phrase,
to (1.0) assuming all the words are completely unrelated.

A phrase can not possibly occur more than its individual words.
Therefor the highest frequent possible for a phrase is the lowest frequency of its component words.
The phrase "new york" can not be more common than the word "york" which it contains.
The estimate uses the lowest word frequency in the phrase as a maximum value.

For a minimum we combined the probability of the component words occurring individually using probability math.
`P(A and B and C) ≈ P(A) * P(B) * P(C)` where P(n) is a proportion between 1 and 0.
This value assumes the words occur like dice rolls which are independent and unrelated.

Then the estimate takes a value somewhere between the minimum and maximum possible frequencies.
The default bias is 0.5. Which picks a point halfway between the two.
This will over estimate uncorrelated phrases and under estimate correlated phrases.

I'm inclined to think that most phrases provided by a user are likely to be more correlated than not.

We can use the lowest frequency as the top of your estimate.




where the word never occurs outside of the phrase.

"new york" can not possibly occur more frequently then its parts.




