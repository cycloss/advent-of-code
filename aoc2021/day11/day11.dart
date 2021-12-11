import 'dart:convert';
import 'dart:io';

import 'objects.dart';

void main() async {
  Stream<String> lines = File("day11.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());

  var rows = <List<int>>[];

  await for (var line in lines) {
    rows.add(parseLine(line));
  }

  int flashes = OctoSim(rows: rows).iterateOctos(100);
  int firstSimultaneous = OctoSim(rows: rows).findFirstSimultaneousFlash();
  print("Part 1 Solution: $flashes");
  print("Part 2 Solution: $firstSimultaneous");
}

List<int> parseLine(String line) {
  return line.split("").map((e) => int.parse(e)).toList();
}
