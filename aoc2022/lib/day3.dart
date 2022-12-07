import 'dart:convert';
import 'dart:io';

void main() async {
  var lines = await File("inputs/day3")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter())
      .toList();
  var totalScore1 = 0;
  for (var line in lines) {
    var chars = line.split("");
    var section1 = chars.sublist(0, chars.length ~/ 2);
    var section2 = chars.sublist(chars.length ~/ 2, chars.length);

    var common = section1.toSet().intersection(section2.toSet()).first;
    totalScore1 += getPriorityScore(common);
  }
  print(totalScore1);

  var totalScore2 = 0;
  var bags = <Set<String>>[];
  for (var i = 0; i < lines.length; i++) {
    bags.add(lines[i].split("").toSet());
    if ((i + 1) % 3 == 0) {
      var common = bags[0].intersection(bags[1]).intersection(bags[2]).first;
      bags.clear();
      totalScore2 += getPriorityScore(common);
    }
  }
  print(totalScore2);
}

int getPriorityScore(String char) {
  var ascii = char.codeUnits[0];
  if (ascii < 91) {
    return ascii - 64 + 26;
  }
  return ascii - 96;
}
