# Day 12 Strategy

## Parsing

- Go through the cave pair lines twice
    - First time through:
        - create a hashmap with node name as key, and node as value
        - for each cave pair, look up each cave in map, if doesn't exist, create a node for it, initialising the node with its name, smallCave, and empty connecting nodes list
    - Second time through:
        - for each cave pair look up the nodes in the map, and link them.
    - return the map to solution object

## Algorithm

- Initialise a route with start node
- Start recursion with start node, recurse until all valid routes traversed, returning a list of routes
- For each node connected to node
    - Copy the route, and add the node to it
    - if the node is finish node, return the route
    - if the node is small node and already visited, return nil
    - recurse for each next potential node
- Count the list of routes
