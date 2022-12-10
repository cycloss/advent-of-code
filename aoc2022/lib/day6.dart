import 'dart:io';

void day6() {
  // detect start of packet marker
  // start of packet is sequence of 4 chars that are different
  // how many chars to process before first start of packet marker detected?
  var data = File("inputs/day6").readAsLinesSync()[0].split("");
  print("Part 1:");
  findMarker(data, 4);
  print("Part 2:");
  findMarker(data, 14);
}

void findMarker(List<String> data, int windowLength) {
  for (var i = 0; i < data.length - windowLength; i++) {
    var window = data.sublist(i, i + windowLength);
    var windowSet = window.toSet();
    if (windowSet.length >= windowLength) {
      print(i + windowLength);
      return;
    }
  }
}
