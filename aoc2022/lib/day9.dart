import 'dart:core';
import 'dart:io';
import 'dart:math';

import 'package:aoc2022/tuple.dart';

class Point extends Tuple<int, int> {
  Point(super.x, super.y);
  Point.copy(Point p) : super(p.x, p.y);

  void addV2(Vector2 v) {
    x += v.x;
    y += v.y;
  }

  void alignTo(Point p) {
    x = p.x;
    y = p.y;
  }

  int simpleDistanceFrom(Point p) =>
      max<int>((x - p.x).abs(), (y - p.y).abs()).abs();

  @override
  bool operator ==(Object other) =>
      other is Point &&
      other.runtimeType == runtimeType &&
      other.x == x &&
      other.y == y;

  @override
  int get hashCode => Object.hash(x, y);
}

class Vector2 extends Tuple<int, int> {
  factory Vector2.normalisedFromDirection(String direction) =>
      ordinals[direction]!;

  Vector2(super.x, super.y);

  static final ordinals = {
    "U": Vector2(0, 1),
    "D": Vector2(0, -1),
    "L": Vector2(-1, 0),
    "R": Vector2(1, 0),
  };
}

final origin = Point(0, 0);

void day9() {
  var lines = File("inputs/day9").readAsLinesSync();
  // points tail has visited

  var visited = <Point>{origin};
  var head = Point(0, 0), prevHead = Point(0, 0), tail = Point(0, 0);

  for (var line in lines) {
    var split = line.split(" ");
    var dir = split[0];
    var v = Vector2.normalisedFromDirection(dir);
    var steps = int.parse(split[1]);
    for (var i = 0; i < steps; i++) {
      // add v to head and check where tail should be
      prevHead = Point.copy(head);
      head.addV2(v);
      // check where tail should now be
      var dist = head.simpleDistanceFrom(tail);
      if (dist > 1) {
        tail.alignTo(prevHead);
      }
      // add tail now pos to visited set
      var copy = Point.copy(tail);
      visited.add(copy);
    }
  }

  print(visited.length);
}
