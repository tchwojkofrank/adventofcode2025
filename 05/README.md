# adventofcode2025
## Part 1
The input is a set of intervals, and a set of integers.
Count how many integers are in at least one interval
## Part 2
Ignore the set of integers.
Measure the size of the union of the intervals.
In this case we actually need to remove any overlaps to make it easy to count.
So for each interval, subtract any overlap with the existing intervals, and add the result of that subtraction to the list.
Now all the intervals are guaranteed to be discrete, so we can measure each one individually.