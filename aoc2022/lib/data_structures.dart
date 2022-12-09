import 'dart:collection';

class Tuple<T, U> {
  T x;
  U y;
  Tuple(this.x, this.y);

  @override
  String toString() => 'x: $x, y: $y';

  @override
  bool operator ==(Object other) =>
      other is Tuple &&
      other.runtimeType == runtimeType &&
      other.x == x &&
      other.y == y;

  @override
  int get hashCode => Object.hash(x, y);
}

class Vector2 extends Tuple<int, int> {
  factory Vector2.fromNormalisedDir(String direction) => ordinals[direction]!;

  Vector2(super.x, super.y);

  static final ordinals = {
    "U": Vector2(0, 1),
    "D": Vector2(0, -1),
    "L": Vector2(-1, 0),
    "R": Vector2(1, 0),
  };
}

class Stack<T> extends ListQueue<T> {
  void pushAll(Iterable<T> v) {
    addAll(v);
  }

  List<T> popCount(int count) {
    var out = <T>[];
    for (var i = 0; i < count; i++) {
      if (isNotEmpty) {
        out.insert(0, removeLast());
      }
    }
    return out;
  }

  void push(T v) => add(v);

  T pop() => removeLast();
}

class TreeNode<T> {
  T value;

  TreeNode(this.value);

  List<TreeNode<T>> children = [];

  void add(TreeNode<T> child) {
    children.add(child);
  }
}
