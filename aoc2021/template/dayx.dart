import 'dart:convert';
import 'dart:io';

import 'parser.dart';

void main() async {
  var lines = await File("day12.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter())
      .toList();

  var solution = Solution(lines);

  print("Part 1 Solution: ${solution.part1()}");
  print("Part 2 Solution: ${solution.part2()}");
}

class Solution {
  final List<String> lines;
  Solution(List<String> lines) : lines = Parser(lines).processLines();

  int part1() {
    return lines.length;
  }

  int part2() {
    return lines.length;
  }
}
