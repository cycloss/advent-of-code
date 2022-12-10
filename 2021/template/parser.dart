class Parser {
  final List<String> lines;
  List<String> out = [];
  Parser(this.lines);

  List<String> processLines() {
    out = lines;
    return out;
  }
}
