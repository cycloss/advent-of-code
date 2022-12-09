import 'dart:io';

void dayn() {
  var lines = File("inputs/dayn").readAsLinesSync();
  // var chars = File("/inputs/dayn").readAsBytesSync();
  for (var line in lines) {
    print(line);
  }
}
