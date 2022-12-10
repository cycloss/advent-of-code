import 'dart:core';
import 'dart:io';
import 'dart:math';

import 'package:2022/data_structures.dart';

class Point extends Tuple<int, int> {
  Point(super.x, super.y);
  Point.copy(Point p) : super(p.x, p.y);
  Point.origin() : super(0, 0);

  void addV2(Vector2 v) {
    x += v.x;
    y += v.y;
  }

  int simpleDistanceFrom(Point p) =>
      max<int>((x - p.x).abs(), (y - p.y).abs()).abs();

  Vector2 normalisedVectorTo(Point p) {
    var xDir = p.x - x;
    var yDir = p.y - y;
    xDir = max(-1, min(1, xDir));
    yDir = max(-1, min(1, yDir));
    return Vector2(xDir, yDir);
  }
}

void day9() {
  var lines = File("inputs/day9").readAsLinesSync();
  print("Part 1:");
  calculateTailPath(2, lines);
  print("Part 2:");
  calculateTailPath(10, lines);
}

void calculateTailPath(int ropeKnotCount, List<String> lines) {
  var tailVisited = <Point>{Point.origin()};
  var ropeKnots = List.generate(ropeKnotCount, (_) => Point.origin());

  for (var line in lines) {
    var split = line.split(" ");
    var dir = split[0];
    var v = Vector2.fromNormalisedDir(dir);

    for (var steps = int.parse(split[1]); steps > 0; steps--) {
      for (var i = 0; i < ropeKnotCount; i++) {
        var knot = ropeKnots[i];
        if (i == 0) {
          // knot is head
          knot.addV2(v);
        } else {
          var prevKnot = ropeKnots[i - 1];
          var dist = knot.simpleDistanceFrom(prevKnot);
          if (dist > 1) {
            var v2 = knot.normalisedVectorTo(prevKnot);
            knot.addV2(v2);
            if (i == ropeKnots.length - 1) {
              // is tail
              var copy = Point.copy(knot);
              tailVisited.add(copy);
            }
          }
        }
      }
    }
  }

  print(tailVisited.length);
}
