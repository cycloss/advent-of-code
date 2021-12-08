# Day 8 Strategy

## Algorithm With Example

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab

- ab is 1, dab is 7, eafb is 4, acedgfb is 8
- 6 segs: cefabd cdfgeb cagedb
- 5 segs: cdfbe gcdfa fbcad
- top seg: not common to 1 and 7: d
- bottom seg: common to 6 segs: c e b d, not in 4 c, d, not top seg: c
- bottom left seg: not in 1, 4, 7: c? g?: g
- top right: in 1 but only 2 of 6 segs: a
- bottom right: b
- middle: common to 5 segs c, d, f, not known: f
- top left remaining not known: e
- **correct

- create sets for known 1, 4, 7 and 8
- create array of sets for 5 and 6 sets

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

  aaaa
 b    c  
 b    c  
  dddd
 e    f  
 e    f  
  gggg

  // unique: 1 cf, 4 bcdf, 7 acf, 8 a-g
// non unique: [6seg 0, 6, 9], [5seg 2,3, 5]
// top segment is one not common to 1 and 7
// bottom seg will be present in all the 6 segs, but not in 4, and isn't the top seg
// bottom left seg is only one present in 8 but not in 1 4 and 7 and isn't the bottom seg
// top right is present in 1, and only 2 of the 6 segs
// now know bottom right as is not top right of 1
// middle seg is common to all 5 segs and not known
// top left seg is only one not known in 4
