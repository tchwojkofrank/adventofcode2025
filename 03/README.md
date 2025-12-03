# adventofcode2025
## Part 1
Find the largest possible two digit number in reading order in each line.
Check the 10s digit first, from 0..len-1 then the ones from the current index to the length of the string.
## Part 2
Find the largest possible 12 digit number in reading order in each line.
Start with the 10^11 digit, and find the largest up to index len-11
Then from that index, find the largest up to index len-10 for the 10^10 digit
Repeat until we get to the ones.