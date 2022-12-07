import 'dart:convert';
import 'dart:io';

enum Outcome { win, loose, draw }

final translationMap = {
  "A": "R",
  "X": "R",
  "B": "P",
  "Y": "P",
  "C": "S",
  "Z": "S"
};

final choiceScoreMap = {
  "R": 1,
  "P": 2,
  "S": 3,
};
final outcomesScoreMap = {
  Outcome.loose: 0,
  Outcome.draw: 3,
  Outcome.win: 6,
};

final rockMap = {"R": Outcome.draw, "P": Outcome.loose, "S": Outcome.win};
final paperMap = {"P": Outcome.draw, "R": Outcome.win, "S": Outcome.loose};
final scissorsMap = {"S": Outcome.draw, "R": Outcome.loose, "P": Outcome.win};
final winMap = {"R": rockMap, "P": paperMap, "S": scissorsMap};

final commandMap = {"X": Outcome.loose, "Y": Outcome.draw, "Z": Outcome.win};
final part2Map = {"R": rockMap2, "P": paperMap2, "S": scissorsMap2};
// if they have chosen rock, and you have command, then what do you need
final rockMap2 = {Outcome.draw: "R", Outcome.loose: "S", Outcome.win: "P"};
final paperMap2 = {Outcome.draw: "P", Outcome.loose: "R", Outcome.win: "S"};
final scissorsMap2 = {Outcome.draw: "S", Outcome.loose: "P", Outcome.win: "R"};

void main() async {
  Stream<String> lines = File("inputs/day2")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());

  // 0 lost, 3 draw, 6 win
  var totalScorePart1 = 0;
  var totalScorePart2 = 0;

  await for (var line in lines) {
    var commands = line.split(" ");
    var p1Choice = translationMap[commands[0]];
    // part1
    var p2Choice = translationMap[commands[1]];
    var p2ChoiceScore = choiceScoreMap[p2Choice]!;
    totalScorePart1 += p2ChoiceScore;
    var p2Outcome = winMap[p2Choice]![p1Choice]!;
    var p2OutcomeScore = outcomesScoreMap[p2Outcome]!;
    totalScorePart1 += p2OutcomeScore;
    // part2
    // x loose, y draw, z win
    var p2Command = commandMap[commands[1]]!;
    // calculate outcome score
    var p2OutcomeScore2 = outcomesScoreMap[p2Command]!;
    totalScorePart2 += p2OutcomeScore2;
    // calculate choice score
    var p1ChoiceMap2 = part2Map[p1Choice]!;

    var p2Choice2 = p1ChoiceMap2[p2Command]!;
    var p2ChoiceScore2 = choiceScoreMap[p2Choice2]!;
    totalScorePart2 += p2ChoiceScore2;
  }
  print("part1: $totalScorePart1, part2: $totalScorePart2");
}
