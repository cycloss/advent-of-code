import 'dart:convert';
import 'dart:io';

void main() async {
  Stream<String> lines = File("day9.txt")
      .openRead()
      .transform(Utf8Decoder())
      .transform(LineSplitter());

  await for (var line in lines) {
    print(line);
  }

  print("Part 1 Solution: ${0}");
  print("Part 2 Solution: ${0}");
}
