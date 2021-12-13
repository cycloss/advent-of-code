import 'dart:convert';
import 'dart:io';

import 'parser.dart';

void main() async {
  var lines = await File("day13.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter())
      .toList();

  var solution = Solution(lines);
  print("Part 1 Solution: ${solution.part1()}");

  var solution2 = Solution(lines);
  solution2.part2();
}

class Solution {
  final Problem problem;

  Solution(List<String> lines) : problem = Parser(lines).processLines();

  int part1() {
    for (var ins in problem.instructions) {
      if (ins.axis == 'x') {
        problem.grid = foldXAxis(problem.grid, ins.index);
        break;
      } else {
        problem.grid = foldYAxis(problem.grid, ins.index);
        break;
      }
    }
    return countVisibleDots(problem.grid);
  }

  List<List<bool>> foldXAxis(List<List<bool>> grid, int index) {
    for (var c = grid[0].length - 1; c > index; c--) {
      var rc = (grid[0].length - 1) - c;
      for (var j = 0; j < grid.length; j++) {
        var b = grid[j][c];
        if (b) {
          grid[j][rc] = true;
        }
      }
    }
    var newRows = <List<bool>>[];

    for (var c = 0; c < index; c++) {
      for (var j = 0; j < grid.length; j++) {
        if (newRows.length - 1 < j) {
          newRows.add([]);
        }
        newRows[j].add(grid[j][c]);
      }
    }
    return newRows;
  }

  List<List<bool>> foldYAxis(List<List<bool>> grid, int index) {
    for (var i = grid.length - 1; i > index; i--) {
      var ri = (grid.length - 1) - i;
      var row = grid[i];
      var rrow = grid[ri];
      for (var j = 0; j < row.length; j++) {
        if (row[j]) {
          rrow[j] = true;
        }
      }
    }
    var newRows = <List<bool>>[];
    for (var i = 0; i < index; i++) {
      newRows.add(grid[i]);
    }
    return newRows;
  }

  int countVisibleDots(List<List<bool>> grid) {
    return grid.fold(0,
        (tot, row) => tot + row.fold(0, (tot2, dot) => tot2 + (dot ? 1 : 0)));
  }

  void part2() {
    for (var ins in problem.instructions) {
      if (ins.axis == 'x') {
        problem.grid = foldXAxis(problem.grid, ins.index);
      } else {
        problem.grid = foldYAxis(problem.grid, ins.index);
      }
    }
    for (var row in problem.grid) {
      for (var b in row) {
        stdout.write(b ? "#" : " ");
      }
      stdout.write("\n");
    }
  }
}
