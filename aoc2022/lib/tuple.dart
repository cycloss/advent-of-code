class Tuple<T, U> {
  T x;
  U y;
  Tuple(this.x, this.y);

  @override
  String toString() => 'x: $x, y: $y';
}
