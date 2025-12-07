# adventofcode2025 Day 7
## Part 1
Splitting beams
Beams travel vertically down, and sometimes hit a splitter.
Count how many times a beam is split.
The grid provided helpfully has blank lines to keep track of beams.
Iterate over each line and figure out which columns have a beam.
## Part 2
Now figure out how many possible paths there are.
I tried a brute force depth first search, and that was taking far too long O(2^n).
But since subtrees always have the same value, we can just keep a ledger to drastically reduce the size of the tree traversal.

## Performance
```
Result: 1602
Running time: 36.042µs
Result 2: 135656430050438
Running time: 370.833µs
```