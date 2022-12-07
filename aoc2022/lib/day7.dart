import 'dart:io';

var allTotal = 0;

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

  int totalSize() {
    var total = size;
    for (var child in children) {
      var childTotal = child.totalSize();
      total += childTotal;
    }
    return total;
  }

  int childNodeForAtMost(int bytes) {
    if (children.isEmpty) {
      // if file
      return size;
    }
    // directory
    var total = 0;
    for (var child in children) {
      var childTotal = child.childNodeForAtMost(bytes);
      total += childTotal;
    }
    if (total <= bytes) {
      allTotal += total;
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

  FSTraverser(this.lines);

  String get currentLine => lines[currentLineNo];
  List<String> get currentLineTokens => currentLine.split(" ");
  bool get currentLineIsCommand => currentLineTokens[0] == "\$";

  int bytes = 0;

  void findDirectoriesOfAtMost(int bytes) {
    this.bytes = bytes;
    currentLineNo = 0;
    buildDirTree();
    // headNode.childNodeForAtMost(bytes);
    var smallDirsTotal = findSmallDirsTotal(headNode);
    print(smallDirsTotal);
  }

  void buildDirTree() {
    while (currentLineNo < lines.length) {
      consumeCommand();
    }
  }

  int findSmallDirsTotal(FSNode node) {
    // only visit directories
    if (node.children.isEmpty) {
      return 0;
    }
    var nodeSize = node.totalSize();
    if (nodeSize <= bytes) {
      return nodeSize;
    }
    var total = 0;
    for (var childNode in node.children) {
      total += findSmallDirsTotal(childNode);
    }
    return total;
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
}

Future<void> day7() async {
  var lines = await File("inputs/day7").readAsLines();
  // use a doubly linked list to traverse the structure and create a tree
  // as going back up, keep a total of the size of dirs you have visited in the top node
  var traverser = FSTraverser(lines);
  traverser.findDirectoriesOfAtMost(100000);
}
