import 'dart:convert';
import 'dart:io';

void main() async {
  Stream<String> lines = File("inputs/day1")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());

  var top3 = [0, 0, 0];
  var currentCals = 0;
  await for (var line in lines) {
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
  print(top3.reduce((value, element) => value + element));
}
