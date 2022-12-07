import 'dart:convert';
import 'dart:io';

import 'parser.dart';

void main() async {
  var lines = await File("day14.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter())
      .toList();

  var solution = Solution(lines);
  print("Part 1 Solution: ${solution.solve(10)}");
  var solution2 = Solution(lines);
  print("Part 2 Solution: ${solution2.solve(40)}");
}

class Solution {
  final Problem problem;
  Solution(List<String> lines) : problem = Parser(lines).processLines();

  int solve(int iterations) {
    for (var i = 0; i < iterations; i++) {
      iterate();
      print(i);
    }
    return calculateAnswer(problem.polymer);
  }

  void iterate() {
    var polymer = problem.polymer;
    var rules = problem.rules;
    for (Node? current = polymer.head!;
        current?.next != null;
        current = current.next) {
      if (current == null) {
        break;
      }
      var next = current.next!;
      var cv = current.value;
      var nv = next.value;
      var pattern = Pattern(cv, nv);
      if (rules.containsKey(pattern)) {
        var newNode = Node(rules[pattern]!);
        newNode.next = next;
        current.next = newNode;
        current = newNode;
      }
    }
  }

  int calculateAnswer(LinkedList polymer) {
    var counts = <String, int>{};
    for (var node = polymer.head; node != null; node = node.next) {
      counts.update(node.value, (value) => ++value, ifAbsent: () => 1);
    }
    var smallest = 1 << 63 - 1;
    var largest = 0;
    for (var e in counts.entries) {
      if (e.value < smallest) {
        smallest = e.value;
      }
      if (e.value > largest) {
        largest = e.value;
      }
    }
    return largest - smallest;
  }
}
