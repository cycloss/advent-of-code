import 'dart:convert';
import 'dart:io';

void main() async {
  Stream<String> lines = File("day10.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());

  var totalSyntaxScore = 0;
  var autoCompleteScores = <int>[];
  await for (var line in lines) {
    var result = findSyntaxScore(line);
    totalSyntaxScore += result.syntaxScore;
    if (result.autoCompleteScore != 0) {
      autoCompleteScores.add(result.autoCompleteScore);
    }
  }
  autoCompleteScores.sort();
  var middleIndex = (autoCompleteScores.length / 2).round() - 1;
  var middleScore = autoCompleteScores[middleIndex];
  print("Part 1 Solution: $totalSyntaxScore");
  print("Part 2 Solution: $middleScore");
}

class ScoreResult {
  final int syntaxScore;
  final int autoCompleteScore;

  ScoreResult(this.syntaxScore, this.autoCompleteScore);
}

var opening = <String>{'(', '[', '<', '{'};
var closing = <String>{')', ']', '>', '}'};

var pairs = <String, String>{')': '(', ']': '[', '>': '<', '}': '{'};
var corruptScores = <String, int>{')': 3, ']': 57, '>': 25137, '}': 1197};

var closingPairs = <String, String>{'(': ')', '[': ']', '<': '>', '{': '}'};
var completeScores = <String, int>{')': 1, ']': 2, '>': 4, '}': 3};

ScoreResult findSyntaxScore(String line) {
  var brackets = line.split("");
  var bracketStack = <String>[];

  for (var b in brackets) {
    if (bracketStack.isEmpty || opening.contains(b)) {
      bracketStack.add(b);
    } else {
      var top = bracketStack.last;
      if (pairs[b] == top) {
        bracketStack.removeLast();
      } else {
        return ScoreResult(corruptScores[b]!, 0);
      }
    }
    // stop at first ilegal closing
  }
  // not corrupted if got to here
  var completeTotal = 0;
  while (bracketStack.isNotEmpty) {
    var top = bracketStack.removeLast();
    var closer = closingPairs[top];
    var score = completeScores[closer]!;
    completeTotal *= 5;
    completeTotal += score;
  }
  return ScoreResult(0, completeTotal);
}
