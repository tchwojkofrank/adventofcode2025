# adventofcode2025
## Part 1
Read the numbers in each column, apply the operator at the bottom to the whole list.
Sum the results
## Part 2
The numbers in the original input are vertical, rather than horizontal, and read right to left.
Read the characters into a grid.
Transpose the mirror of the grid.
This gets the columns from the input from right to left as rows in the new grid.
Each row is now a number, optionally ending with an operator.
After an operator the next row is blank.
For each row, get the value and append into the argument array.
If the row ends in an operator, apply the operator to the arguments in the array, and then add to the total, and clear the array. Skip the next line which is blank.
