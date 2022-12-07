import 'dart:collection';

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

  void push(T v) {
    add(v);
  }
}
