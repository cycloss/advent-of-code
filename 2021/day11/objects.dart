class Octo {
  int x, y, value;
  Octo(this.x, this.y, this.value);

  @override
  int get hashCode => Object.hashAll([y, x]);

  @override
  bool operator ==(Object other) {
    return other is Octo &&
        other.runtimeType == runtimeType &&
        other.y == y &&
        other.x == x;
  }

  bool increment() {
    value++;
    if (value > 9) {
      value = 0;
      return true;
    }
    return false;
  }
}

class OctoSim {
  late final List<List<Octo>> octos;
  int flashes = 0;

  int get octoCount =>
      octos.fold(0, (tot, row) => tot + row.fold(0, (prev, _) => prev + 1));

  OctoSim({
    required rows,
  }) {
    initOctos(rows);
  }

  void initOctos(List<List<int>> rows) {
    List<List<Octo>> octos = [];
    for (var y = 0; y < rows.length; y++) {
      var row = rows[y];
      var octoRow = <Octo>[];
      for (var x = 0; x < row.length; x++) {
        var value = row[x];
        var octo = Octo(x, y, value);
        octoRow.add(octo);
      }
      octos.add(octoRow);
    }
    this.octos = octos;
  }

  int iterateOctos(int iterations) {
    for (var i = 0; i < iterations; i++) {
      iterate();
    }
    return flashes;
  }

  int findFirstSimultaneousFlash() {
    int iterations = 0;
    while (true) {
      iterations++;
      var currentFlashCount = iterate();
      if (currentFlashCount >= octoCount) {
        break;
      }
    }
    return iterations;
  }

  int iterate() {
    var flashed = <Octo>{};
    var currentFlashCount = 0;
    for (var y = 0; y < octos.length; y++) {
      var octoRow = octos[y];
      for (var x = 0; x < octoRow.length; x++) {
        var octo = octoRow[x];
        currentFlashCount += processOcto(octo, flashed);
      }
    }
    return currentFlashCount;
  }

  int processOcto(Octo octo, Set<Octo> flashed) {
    if (flashed.contains(octo)) {
      return 0;
    }
    var didFlash = octo.increment();
    if (didFlash) {
      flashed.add(octo);
      flashes++;
      return 1 +
          findValidNeighbours(octo)
              .fold(0, (total, o) => total + processOcto(o, flashed));
    }
    return 0;
  }

  List<Octo> findValidNeighbours(Octo octo) {
    var valids = <Octo>[];
    var x = octo.x;
    var y = octo.y;
    var upOk = y > 0;
    var downOk = y < octos.length - 1;
    var leftOk = x > 0;
    var rightOk = x < octos[y].length - 1;
    if (upOk) {
      valids.add(octos[y - 1][x]);
      if (rightOk) {
        valids.add(octos[y - 1][x + 1]);
      }
      if (leftOk) {
        valids.add(octos[y - 1][x - 1]);
      }
    }
    if (rightOk) {
      valids.add(octos[y][x + 1]);
    }
    if (downOk) {
      valids.add(octos[y + 1][x]);
      if (rightOk) {
        valids.add(octos[y + 1][x + 1]);
      }
      if (leftOk) {
        valids.add(octos[y + 1][x - 1]);
      }
    }
    if (leftOk) {
      valids.add(octos[y][x - 1]);
    }
    return valids;
  }
}
