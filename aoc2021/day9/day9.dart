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

  var riskLevel = 0;

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
        riskLevel += 1 + num;
      }
    }
  }

  print("Part 1 Solution: $riskLevel");
}
