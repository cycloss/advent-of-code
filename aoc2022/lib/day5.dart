import 'dart:collection';
import 'dart:io';

class Instruction {
  final int countToMove, fromStack, toStack;
  Instruction._(this.countToMove, this.fromStack, this.toStack);

  @override
  String toString() =>
      "move $countToMove from ${fromStack + 1} to ${toStack + 1}";

  factory Instruction.fromLine(String line) {
    var s = line.split(" ");
    return Instruction._(
        int.parse(s[1]), int.parse(s[3]) - 1, int.parse(s[5]) - 1);
  }
}

class CrateGame {
  final List<Stack> crateStacks;

  CrateGame(this.crateStacks);
  void apply(Instruction i) {
    var movingCrates = crateStacks[i.fromStack].pop(i.countToMove);
    crateStacks[i.toStack].push(movingCrates);
  }

  void printTopCrates() {
    for (var stack in crateStacks) {
      if (stack.isNotEmpty) {
        stdout.write(stack.last);
      } else {
        print(" ");
      }
    }
    print("");
  }
}

class Stack extends ListQueue<String> {
  final bool preserveOrder;

  Stack(this.preserveOrder);

  void push(Iterable<String> v) {
    addAll(v);
  }

  List<String> pop(int count) {
    var out = <String>[];
    for (var i = 0; i < count; i++) {
      if (isNotEmpty) {
        if (preserveOrder) {
          out.insert(0, removeLast());
        } else {
          out.add(removeLast());
        }
      }
    }
    return out;
  }
}

void day5() async {
  var lines = File("inputs/day5").readAsLinesSync();
  // first 8 lines parse 9 crate stacks
  var stacks1 = parseInitialCrateStacks(lines.sublist(0, 8), false);
  var stacks2 = parseInitialCrateStacks(lines.sublist(0, 8), true);

  var cg1 = CrateGame(stacks1);
  var cg2 = CrateGame(stacks2);
  for (var i = 10; i < lines.length; i++) {
    var instruction = Instruction.fromLine(lines[i]);
    cg2.apply(instruction);
    cg1.apply(instruction);
  }
  print("Part 1:");
  cg1.printTopCrates();
  print("Part 2:");
  cg2.printTopCrates();
}

List<Stack> parseInitialCrateStacks(List<String> lines, bool preserveOrder) {
  var stacks = List.generate(9, (_) => Stack(preserveOrder));
  for (var line in lines) {
    // 1,5,9
    var chars = line.split("");
    // parse every crate on the line, prepending to stack
    for (var i = 0; i < 9; i++) {
      var cratePosition = (4 * i) + 1;
      var char = chars[cratePosition];
      if (char != " ") {
        stacks[i].addFirst(char);
      }
    }
  }
  return stacks;
}
