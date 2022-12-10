import 'dart:io';

// File system node
class FSNode {
  // size is current size of dir
  final String name;
  int size;
  FSNode? parentNode;
  List<FSNode> children = [];
  // name (value in super) is node name
  // children is node children
  // node with no children is either an empty file or directory

  FSNode(this.name, this.size, this.parentNode);

  bool get isDirectory => size == 0;

  int totalSize() {
    if (!isDirectory) {
      return size;
    }
    var total = 0;
    for (var child in children) {
      var childTotal = child.totalSize();
      total += childTotal;
    }
    return total;
  }

  void add(FSNode child) {
    children.add(child);
  }
}

class FSTraverser {
  final List<String> lines;
  int currentLineNo = 0;
  late FSNode headNode;
  late FSNode currentNode;

  FSTraverser(this.lines) {
    buildDirTree();
  }

  String get currentLine => lines[currentLineNo];
  List<String> get currentLineTokens => currentLine.split(" ");
  bool get currentLineIsCommand => currentLineTokens[0] == "\$";

  int totalUnderMaxBytes = 0;

  void buildDirTree() {
    while (currentLineNo < lines.length) {
      consumeCommand();
    }
  }

  // after a consume command, the currentLine will be the next one processed
  void consumeCommand() {
    if (!currentLineIsCommand) {
      print('not a command: $currentLine');
      exit(1);
    }
    switch (currentLineTokens[1]) {
      case "cd":
        handleCd(currentLineTokens[2]);
        break;
      case "ls":
        handleLs();
        break;
      default:
        print('unknown command: $currentLine');
        exit(1);
    }
  }

  void handleCd(String dirName) {
    switch (dirName) {
      case "/":
        // get sizes of all nodes underneath and add to current node
        headNode = currentNode = FSNode("/", 0, null);
        break;
      case "..":
        currentNode = currentNode.parentNode!;
        break;
      default:
        var childNode = FSNode(dirName, 0, currentNode);
        currentNode.add(childNode);
        currentNode = childNode;
    }
    currentLineNo++;
  }

  void handleLs() {
    currentLineNo++;
    for (; currentLineNo < lines.length; currentLineNo++) {
      if (currentLineIsCommand) {
        return;
      }
      if (!(currentLineTokens[0] == "dir")) {
        var val = int.parse(currentLineTokens[0]);
        var childNode = FSNode(currentLineTokens[1], val, currentNode);
        currentNode.add(childNode);
      }
    }
  }

  void findDirectoriesOfAtMost(int maxBytes) {
    findSmallSubDirs(headNode, maxBytes);
    print("Part 1:");
    print('Total of dirs under $maxBytes: $totalUnderMaxBytes');
  }

  void findSmallSubDirs(FSNode node, int maxBytes) {
    if (!node.isDirectory) {
      return;
    }
    for (var child in node.children) {
      findSmallSubDirs(child, maxBytes);
    }
    var nodeSize = node.totalSize();
    if (nodeSize <= maxBytes) {
      totalUnderMaxBytes += nodeSize;
    }
  }

  int currentSmallest = 999999999;

  void findDirToDeleteToFreeSpace() {
    var rootDirSize = headNode.totalSize();
    var spaceToFree = 30000000 - (70000000 - rootDirSize);
    findSmallDirToFreeSpace(headNode, spaceToFree);
    print("Part 2:");
    print('Smallest dir to free up $spaceToFree: $currentSmallest');
  }

  void findSmallDirToFreeSpace(FSNode node, int spaceToFree) {
    if (!node.isDirectory) {
      return;
    }
    for (var child in node.children) {
      findSmallDirToFreeSpace(child, spaceToFree);
    }
    var nodeSize = node.totalSize();
    if (nodeSize >= spaceToFree && nodeSize < currentSmallest) {
      currentSmallest = nodeSize;
    }
  }
}

void day7() {
  var lines = File("inputs/day7").readAsLinesSync();
  // use a doubly linked list to traverse the structure and create a tree
  // as going back up, keep a total of the size of dirs you have visited in the top node
  var traverser = FSTraverser(lines);
  traverser.findDirectoriesOfAtMost(100000);
  traverser.findDirToDeleteToFreeSpace();
}
