# Day 5 Strategy

## Parsing

- For each line of input
    - find the smallest x and y values
    - create a struct called ventLine with two vector2s start and end
    - put the smallest x,y pair in start, and largest in end
    - create a bool method diagonal on ventLine that returns true if neither x and y values match

## Algorithm

- Create a hashmap, countMap, of map[vector2]int to count vents at each point
- For each ventLine struct
    - If diagonal
        - continue
    - Else
        - Create x and y values with the start vector2s x and y
        - While they are less than the end x and y values
            - create a vector2, currentCoord, out of x and y
            - increment countMap[currentCoord] by 1
            - increment x and y by 1
- return count of where there are >= 2 vents at a given key
