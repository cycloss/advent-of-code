import 'dart:convert';
import 'dart:io';

void main() async {
  var lines = await File("day12.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter())
      .toList();

  var solution = Solution(lines);

  print("Part 1 Solution: ${solution.findRouteCount()}");
  solution.allowOneTimeDouble = true;
  print("Part 2 Solution: ${solution.findRouteCount()}");
}

class Solution {
  final Map<String, Node> nodeMap;
  bool allowOneTimeDouble = false;
  Solution(List<String> lines) : nodeMap = Parser(lines).generateNodeMap();

  int findRouteCount() {
    var emptyRoute =
        Route(nodeRoute: [], visited: {}, oneSmallVisitedTwice: false);
    var routes = findAllRoutesFrom(nodeMap['start']!, emptyRoute);
    // routes.forEach((r) => r.printRoute());
    return routes.length;
  }

  List<Route> findAllRoutesFrom(Node currentNode, Route routeSoFar) {
    var visited = routeSoFar.nodeAlreadyVisited(currentNode);
    var visitedTwice = routeSoFar.oneSmallVisitedTwice;
    if (visited && currentNode.smallCave) {
      if (!allowOneTimeDouble) {
        return [];
      } else {
        if (currentNode.name == 'start') {
          return [];
        }
        if (routeSoFar.oneSmallVisitedTwice) {
          return [];
        } else {
          visitedTwice = true;
        }
      }
    }

    var copy = routeSoFar.copyAndAdd(currentNode, visitedTwice);
    if (currentNode.name == 'end') {
      return [copy];
    }
    var routes = <Route>[];
    for (var node in currentNode.connectingNodes) {
      var nextRoutes = findAllRoutesFrom(node, copy);
      routes.addAll(nextRoutes);
    }
    return routes;
  }
}

class Parser {
  Map<String, Node> nodeMap = {};
  final List<String> lines;

  Parser(this.lines);

  Map<String, Node> generateNodeMap() {
    setupNodeMap();
    linkNodes();
    return nodeMap;
  }

  void setupNodeMap() {
    nodeMap = {};
    for (var line in lines) {
      var pair = line.split("-");
      processCave(pair[0]);
      processCave(pair[1]);
    }
  }

  void processCave(String cave) {
    if (!nodeMap.containsKey(cave)) {
      nodeMap[cave] = Node(smallCave: isSmallCave(cave), name: cave);
    }
  }

  static bool isSmallCave(String cave) {
    return cave == cave.toLowerCase();
  }

  void linkNodes() {
    for (var line in lines) {
      var pair = line.split("-");
      linkPair(pair[0], pair[1]);
    }
  }

  void linkPair(String cave1, String cave2) {
    var node1 = nodeMap[cave1]!;
    var node2 = nodeMap[cave2]!;
    node1.connectingNodes.add(node2);
    node2.connectingNodes.add(node1);
  }
}

class Node {
  String name;
  bool smallCave;
  List<Node> connectingNodes = [];

  Node({
    required this.name,
    required this.smallCave,
  });
}

class Route {
  List<Node> nodeRoute;
  Set<Node> visited;
  bool oneSmallVisitedTwice;
  Route(
      {required this.nodeRoute,
      required this.visited,
      required this.oneSmallVisitedTwice});

  Route copyAndAdd(Node node, bool visitedTwice) {
    var routeCopy = List<Node>.from(nodeRoute);
    routeCopy.add(node);
    var setCopy = Set<Node>.from(visited);
    setCopy.add(node);
    return Route(
        nodeRoute: routeCopy,
        visited: setCopy,
        oneSmallVisitedTwice: visitedTwice);
  }

  bool nodeAlreadyVisited(Node node) => visited.contains(node);

  void printRoute() {
    nodeRoute.forEach((n) {
      stdout.write("${n.name},");
    });
    print("");
  }
}
