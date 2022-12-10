import 'dart:io';

class Tuple {
  final int x, y;
  Tuple(this.x, this.y);

  @override
  String toString() => "x: $x, y: $y";

  bool containsEntirely(Tuple t) {
    return x <= t.x && y >= t.y;
  }

  // x1 <= y2 && y1 <= x2
  bool containsAtAnyOf(Tuple t) {
    return x <= t.y && y >= t.x;
  }
}

void day4() {
  var lines = File("inputs/day4").readAsLinesSync();

  var containsEntirelyCount = 0;
  var containsAnyOfCount = 0;
  for (var line in lines) {
    var sectionPairs = line.split(",");
    var t1 = tupleFromSectionPair(sectionPairs[0]);
    var t2 = tupleFromSectionPair(sectionPairs[1]);
    if (t1.containsEntirely(t2) || t2.containsEntirely(t1)) {
      containsEntirelyCount++;
    }
    if (t1.containsAtAnyOf(t2) || t2.containsAtAnyOf(t1)) {
      containsAnyOfCount++;
    }
  }

  print("Part 1: $containsEntirelyCount");
  print("Part 2: $containsAnyOfCount");
}

Tuple tupleFromSectionPair(String sectionPair) {
  var sections = sectionPair.split("-");
  return Tuple(int.parse(sections[0]), int.parse(sections[1]));
}
