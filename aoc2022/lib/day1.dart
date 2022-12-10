import 'dart:io';

void day1() {
  var lines = File("inputs/day1").readAsLinesSync();

  var top3 = [0, 0, 0];
  var currentCals = 0;
  for (var line in lines) {
    if (line.isEmpty) {
      top3.sort();
      if (currentCals > top3[0]) {
        top3[0] = currentCals;
      }
      currentCals = 0;
      continue;
    }
    currentCals += int.parse(line);
  }
  print("Part 1:");
  print(top3[2]);
  print("Part 2:");
  print(top3.reduce((value, element) => value + element));
}
