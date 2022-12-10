import 'dart:io';

import 'package:2022/wrapped_primitives.dart';

class Tree {
  final int height;
  bool visible;
  Tree(this.height, this.visible);
  List<int> scenicScores = [];
  int get totalScenicScore {
    var scoreTotal = 1;
    for (var score in scenicScores) {
      scoreTotal *= score;
    }
    return scoreTotal;
  }
}

void day8() {
  var lines = File('inputs/day8').readAsLinesSync();
  var grid = generateGrid(lines);
  // tree heights 0 to 9 (0 is a tree)
  // tree visible if all trees from it to edge are shorter than it
  // plan is to start at level 9 and go down to 0
  // do each row forward and back, and each column up and down
  // start all tuples as visible null
  // possibly unnecessary: if index of iteration is 0 (an edge), set visible to true
  // if a visible tree is found before tree of current level, stop and move to next (lower) level as that tree is taller and so nothing after will be visible for the current level
  // if a tree of the current height is found, set it to visible, and all trees after to invisible, until you get to a tree that is visible (will be higher than it, set on a previous iteration)

  findVisibleTrees(grid);
  // printGrid(grid);
  print("Part 1:");
  countVisibleTrees(grid);

  findScenicScores(grid);

  print("Part 2:");
  findMaxScenicScore(grid);
}

void countVisibleTrees(List<List<Tree>> grid) {
  var total = 0;
  for (var row in grid) {
    for (var tree in row) {
      if (tree.visible) {
        total++;
      }
    }
  }
  print(total);
}

void findMaxScenicScore(List<List<Tree>> grid) {
  var max = 0;
  for (var row in grid) {
    for (var tree in row) {
      if (tree.totalScenicScore > max) {
        max = tree.totalScenicScore;
      }
    }
  }
  print(max);
}

void printGrid(List<List<Tree>> grid) {
  for (var row in grid) {
    for (var tree in row) {
      stdout.write('${tree.visible} ');
    }
    print("");
  }
}

void findVisibleTrees(List<List<Tree>> grid) {
  for (var currentLevel = 9; currentLevel >= 0; currentLevel--) {
    for (var row in grid) {
      // iterate row forward for current level
      processVisibilityOfLine(row, currentLevel);
      // iterate row backward for current level
      processVisibilityOfLine(row.reversed, currentLevel);
    }
    for (var column in getGridColumns(grid)) {
      // iterate row forward for current level
      processVisibilityOfLine(column, currentLevel);
      // iterate row backward for current level
      processVisibilityOfLine(column.reversed, currentLevel);
    }
  }
}

List<List<Tree>> getGridColumns(List<List<Tree>> grid) {
  var columns = <List<Tree>>[];
  for (var i = 0; i < grid[0].length; i++) {
    var column = <Tree>[];
    for (var j = 0; j < grid.length; j++) {
      column.add(grid[j][i]);
    }
    columns.add(column);
  }
  return columns;
}

void processVisibilityOfLine(Iterable<Tree> line, int currentLevel) {
  var found = Bool(false);
  for (var tree in line) {
    if (tree.visible) {
      // break as found a tree taller of the same size before found tree of current height
      break;
    }
    processTreeVisibility(tree, found, currentLevel);
  }
}

void processTreeVisibility(Tree tree, Bool found, int currentLevel) {
  if (found.b) {
    // tree must be smaller and therefore not visible as not encountered large visible tree yet, and tree of current level found
    tree.visible = false;
    return;
  }
  if (tree.height == currentLevel) {
    tree.visible = true;
    found.b = true;
  }
}

void findScenicScores(List<List<Tree>> grid) {
  for (var currentLevel = 9; currentLevel >= 0; currentLevel--) {
    for (var row in grid) {
      // iterate row forward for current level
      processScenicScoreOfLine(row, currentLevel);
      // iterate row backward for current level
      processScenicScoreOfLine(row.reversed, currentLevel);
    }
    for (var column in getGridColumns(grid)) {
      // iterate row forward for current level
      processScenicScoreOfLine(column, currentLevel);
      // iterate row backward for current level
      processScenicScoreOfLine(column.reversed, currentLevel);
    }
  }
}

void processScenicScoreOfLine(Iterable<Tree> line, int currentLevel) {
  // set score to 0 as no trees on edge
  var score = 0;
  for (var tree in line) {
    if (tree.height >= currentLevel) {
      if (tree.height == currentLevel) {
        tree.scenicScores.add(score);
      }
      // set score to one as current tree counts
      score = 1;
    } else {
      score++;
    }
  }
}

List<List<Tree>> generateGrid(List<String> lines) {
  var grid = <List<Tree>>[];
  for (var line in lines) {
    var gridLine = <Tree>[];
    for (var char in line.split("")) {
      gridLine.add(Tree(int.parse(char), false));
    }
    grid.add(gridLine);
  }
  return grid;
}
