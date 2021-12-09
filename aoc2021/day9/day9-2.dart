import 'dart:convert';
import 'dart:io';

void main() async {
  Stream<String> lines = File("day9.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());
  var rows = <List<int>>[];
  await for (var line in lines) {
    var split = line.split("");
    var rowNums = <int>[];
    for (var raw in split) {
      rowNums.add(int.parse(raw));
    }
    rows.add(rowNums);
  }

  var basinSizes = <int>[];

  for (var i = 0; i < rows.length; i++) {
    var row = rows[i];
    for (var j = 0; j < row.length; j++) {
      var num = row[j];
      var adjacents = <int>[];

      if (j > 0) {
        var left = row[j - 1];
        adjacents.add(left);
      }
      if (j < row.length - 1) {
        var right = row[j + 1];
        adjacents.add(right);
      }
      if (i > 0) {
        var up = rows[i - 1][j];
        adjacents.add(up);
      }
      if (i < rows.length - 1) {
        var down = rows[i + 1][j];
        adjacents.add(down);
      }

      var lowPoint = true;
      for (var n in adjacents) {
        if (num >= n) {
          lowPoint = false;
        }
      }
      if (lowPoint) {
        var visitedSet = <Vector2>{};
        var size = findBasinSize(rows, visitedSet, Vector2(i, j));
        print(size);
        basinSizes.add(size);
      }
    }
  }
  basinSizes.sort((a, b) => b - a);
  var ans = basinSizes[0] * basinSizes[1] * basinSizes[2];

  print("Part 2 Solution: $ans");
}

class Vector2 {
  int i, j;
  Vector2(this.i, this.j);

  int value(List<List<int>> rows) {
    return rows[i][j];
  }

  @override
  int get hashCode => Object.hashAll([i, j]);

  @override
  bool operator ==(Object other) {
    return other is Vector2 &&
        other.runtimeType == runtimeType &&
        other.i == i &&
        other.j == j;
  }
}

int findBasinSize(List<List<int>> rows, Set<Vector2> visitedSet, Vector2 v2) {
  if (visitedSet.contains(v2)) {
    // already visited
    return 0;
  }
  if (v2.value(rows) == 9) {
    // got to edge
    visitedSet.add(v2);
    return 0;
  }
  visitedSet.add(v2);

  var adjacents = <Vector2>[];
  // add all valid adjacents
  var row = rows[v2.i];
  var i = v2.i;
  var j = v2.j;
  if (j > 0) {
    adjacents.add(Vector2(i, j - 1));
  }
  if (j < row.length - 1) {
    adjacents.add(Vector2(i, j + 1));
  }
  if (i > 0) {
    adjacents.add(Vector2(i - 1, j));
  }
  if (i < rows.length - 1) {
    adjacents.add(Vector2(i + 1, j));
  }
  var total =
      adjacents.fold<int>(0, (t, e) => t + findBasinSize(rows, visitedSet, e));

  return 1 + total;
}
