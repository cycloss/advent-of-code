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
