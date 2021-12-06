# Day 4 Strategy

## Parsing

- For the number sequence, an array
- For the boards, a struct containing:
    - An array for the board itself, containing not numbers, but booleans for if the spot has been marked
    - A map containing the board number as key, and the value as its coordinate on the board
- Also parse in the dimensions of the board, and make them available as globals

## Algorithm

- For each drawn number, go through each board struct
    - For each board struct
        - Check its number set for the drawn number
        - If number not found
            - Skip to next board
        - If found
            - Go to the coordinate on the board, and toggle it to true
            - Then check from top to bottom of the coordinate's column, and left to right of it's row
            - If a complete column or row is found
                - Calculate the score of the board and return that number
- If reached here then fatal error, no solution
