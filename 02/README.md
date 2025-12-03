# Day 2
## Part 1 Finding repeated numbers
Part 1 is to determine the sum of numbers in a set of ranges consists of a single string repeated once.
Some examples of numbers that qualify:
123123
1111

Here I tried to be clever and, instead of iterating through every number in the range, I iterated through the first half (rounded up) of the number in the range.

So if the range was `84-123` I would iterate from 8 to 12, and check if doubling the string was within the range. This greatly reduces the numbers I have to check. There's probably a way to make it even more concise, but I didn't want to think about it that much.
## Part 2 finding more repeated numbers
Part 2 expanded this to include numbers that were any number of repetitions. So the same string could be repeated n times.
Some examples:
111 (three repeats)
569569569569 (four repeats)
8787 (two repeats)
Now my previous optimization gets complicated. I thought I could still be clever, and tried to generalize part 1 with a nested loop, with the outer loop being the number of repeats to check (2, 3, 4 ... len(numberString)). But that got into a whole set of edge cases that became complicated, and I decided to just brute force it. Check each number in the range and determine if it consists of any number of repeated strings.
