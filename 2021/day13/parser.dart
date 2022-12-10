class Parser {
  final List<String> lines;
  List<String> out = [];
  Parser(this.lines);

  int i = 0;

  Problem processLines() {
    var grid = generateGrid();
    var instructions = generateInstructions();
    return Problem(grid: grid, instructions: instructions);
  }

  List<List<bool>> generateGrid() {
    var vectors = <Vector2>[];
    for (; i < lines.length; i++) {
      var line = lines[i];
      if (line == '') {
        i++;
        break;
      }
      var split = line.split(",");
      var vector = Vector2(x: int.parse(split[0]), y: int.parse(split[1]));
      vectors.add(vector);
    }
    var largest = findLargest(vectors);
    var grid = List.generate(
        largest.y + 1, (_) => List.generate(largest.x + 1, (_) => false));
    addPointsToGrid(grid, vectors);
    return grid;
  }

  Vector2 findLargest(List<Vector2> vectors) {
    int maxX = 0;
    int maxY = 0;
    for (var v in vectors) {
      if (v.x > maxX) {
        maxX = v.x;
      }
      if (v.y > maxY) {
        maxY = v.y;
      }
    }
    return Vector2(x: maxX, y: maxY);
  }

  void addPointsToGrid(List<List<bool>> grid, List<Vector2> vectors) {
    for (var v in vectors) {
      grid[v.y][v.x] = true;
    }
  }

  List<Instruction> generateInstructions() {
    var instructions = <Instruction>[];
    for (; i < lines.length; i++) {
      var line = lines[i];
      var split = line.split(" ");
      var instructionStr = split[2];
      var split2 = instructionStr.split("=");
      var instruction = Instruction(split2[0], int.parse(split2[1]));
      instructions.add(instruction);
    }
    return instructions;
  }
}

class Vector2 {
  int x, y;
  Vector2({
    required this.x,
    required this.y,
  });
}

class Problem {
  List<List<bool>> grid;
  final List<Instruction> instructions;
  Problem({
    required this.grid,
    required this.instructions,
  });
}

class Instruction {
  final String axis;
  final int index;

  Instruction(this.axis, this.index);
}
