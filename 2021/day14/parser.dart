class Parser {
  final List<String> lines;
  List<String> out = [];
  Parser(this.lines);
  int index = 0;

  Problem processLines() {
    index = 0;
    var polymer = parsePolymer();
    index += 2;
    var rules = <Pattern, String>{};
    for (; index < lines.length; index++) {
      if (lines[index] == "") {
        continue;
      }
      var raw = lines[index].split(" ");
      var rawPattern = raw[0].split("");
      var pattern = Pattern(rawPattern[0], rawPattern[1]);
      rules[pattern] = raw[2];
    }
    return Problem(polymer, rules);
  }

  LinkedList parsePolymer() {
    var line = lines[index];
    var raw = line.split("");
    return LinkedList(raw);
  }
}

class Problem {
  LinkedList polymer;
  final Map<Pattern, String> rules;
  Problem(this.polymer, this.rules);
}

class Pattern {
  final String first, second;
  Pattern(this.first, this.second);

  @override
  int get hashCode {
    return Object.hashAll([first, second]);
  }

  @override
  bool operator ==(o) {
    return o is Pattern && first == o.first && second == o.second;
  }
}

class LinkedList {
  Node? head;

  LinkedList(List<String> values) {
    late Node current;
    for (var i = 0; i < values.length; i++) {
      if (head == null) {
        head = Node(values[i]);
        current = head!;
      } else {
        var next = Node(values[i]);
        current.next = next;
        current = next;
      }
    }
  }
}

class Node {
  String value;
  Node? next;
  Node(this.value);
}
