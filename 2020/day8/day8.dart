

// void main() async {
//   var ops = await File("day8.txt").readAsString().then((str) => str.split("\n")).then((lines) => lines.map((line) {
//             var matches = RegExp("(\\w+) (\\+|-)(\\d+)").firstMatch(line);
//             return Tuple3<String, int, bool>(matches.group(1), matches.group(2) == "+" ? int.parse(matches.group(3)) : -int.parse(matches.group(3)),false);
//           }).toList());
//   for (var i = 0, acc = 0, tup = ops[i]; !tup.item3; tup.item1 == "acc" ? acc += tup.item2 : tup.item1 == "jmp" ? i += tup.item2 - 1 : 0, ops[i++] = tup.withItem3(true), tup = ops[i], tup.item3 ? print(acc) : 0);
  
// }
