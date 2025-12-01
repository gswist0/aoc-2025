import 'dart:io';

Future<void> main() async {
  var dial = 50;
  var part1 = 0;
  var part2 = 0;
  var file = File("input.txt");
  var lines = await file.readAsLines();
  for (final row in lines){
    var zeros = 0;
    (dial, zeros) = rotate(dial, row);
    if(dial == 0){
      part1++;
    }
    part2 += zeros;
  }
  print(part1);
  print(part2);
}

(int,int) rotate(int currentDial, String row){
  var starting = currentDial;
  var dir = row[0];
  var zeros = 0;
  var count = int.parse(row.substring(1));
  switch(dir){
    case 'L':
      zeros = ((starting - count) / 100).floor().abs();
      currentDial = (starting - count) % 100;
      if(starting == 0){
        zeros--;
      }
      if(currentDial == 0){
        zeros++;
      }
    case 'R':
      currentDial = (starting + count) % 100;
      zeros = ((starting + count) / 100).floor();
  }
  return (currentDial, zeros);
}