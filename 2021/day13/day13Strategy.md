# Day x Strategy

## Parsing

- On first pass:
    - find the largest x and y values
    - when get to break, parse the instructions into struct with axis as string and int as index
- On second pass:
    - init 2d array: array of array of bools with outer of length == largest y and inner length == largest x, filling the whole thing with bools
    - go through each coordinate and set to true in 2d array

## Algorithm

- for each instruction
    - if instruction.axis is x
        - same as y but for columns
    - if instruction.axis is y
        - start on the last row
        - for i = problem.grid.length - 1; i > instruction.index; i--
            - reflectedIndex = problem.grid.length - 1 - i
            - row = problem.grid[i]
            - reflectedRow = problem.grid[refIndex]
            - for j = 0; j < row.length; j++
                - b = row[j]
                - if b:
                    - reflectedRow[j] = true
        - newRows = `List<List<bool>>` = []
        - for i = 0; i < instruction.index; i++
            - copy row to new rows
        - set problem.grid to newRows
