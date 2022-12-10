import 'package:aoc2022/day1.dart';
import 'package:aoc2022/day2.dart';
import 'package:aoc2022/day3.dart';
import 'package:aoc2022/day4.dart';
import 'package:aoc2022/day5.dart';
import 'package:aoc2022/day6.dart';
import 'package:aoc2022/day7.dart';
import 'package:aoc2022/day8.dart';
import 'package:aoc2022/day9.dart';

void main(List<String> arguments) {
  var days = {
    1: day1,
    2: day2,
    3: day3,
    4: day4,
    5: day5,
    6: day6,
    7: day7,
    8: day8,
    9: day9
  };
  if (arguments.isNotEmpty) {
    var dayNum = int.parse(arguments[0]);
    print("Day $dayNum solution:");
    days[dayNum]!();
  } else {
    for (var e in days.entries) {
      print("Day ${e.key} solution:");
      e.value();
      print("----------------------");
    }
  }
}
