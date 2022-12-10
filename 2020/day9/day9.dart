import 'dart:async';
import 'dart:collection';
import 'dart:convert';
import 'dart:io';

void main() {
  //make a queue holding the previous numbers
  //make a set that holds those numbers
  //get the test number
  //iterate through the queue elements and ask if the hashmap contains test number - element
  //if it does, move forward, insert the test number into the hashmap and queue, and remove the last element from the queue, also removing it from the hashmap
  final sw = Stopwatch()..start();
  Future<int> breakingNumber = Solution().findRuleBreakingNumber();
  breakingNumber.then((value) => print(
      "Rule breaking number is: $value\nProgram took: ${sw.elapsedMilliseconds} milliseconds"));
}

class Solution {
  final int queueMax = 25;

  Queue<int> queue = Queue();
  Set<int> numSet = Set();

  Stream<int> getNums() async* {
    Stream<String> lines2 = File("day9.txt")
        .openRead()
        .transform(Utf8Decoder())
        .transform(LineSplitter());

    await for (var line in lines2) {
      int parsed = int.parse(line);
      yield parsed;
    }
  }

  Future<int> findRuleBreakingNumber() async {
    Stream<int> nums = getNums();
    await for (int num in nums) {
      if (queue.length < 25) {
        queue.add(num);
        numSet.add(num);
      } else {
        bool found = findTwoSum(num);
        if (!found) {
          return num;
        }
        int removed = queue.removeFirst();
        numSet.remove(removed);
        queue.add(num);
        numSet.add(num);
      }
    }
    return -1;
  }

  bool findTwoSum(int num) {
    for (int n in queue) {
      int searchNum = num - n;
      if (numSet.contains(searchNum)) {
        return true;
      }
    }
    return false;
  }

  Future<int> findPartTwo(int n) async {
    Queue<int> range = Queue();

    return range.first + range.last;
  }

  Solution() {}
}
