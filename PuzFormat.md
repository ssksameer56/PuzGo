Original Article - `https://code.google.com/archive/p/puz/wikis/FileFormat.wiki`

A Puz File contains following sections
- Header(Size and other data)
- Puzzle Solution
- Current State
- Strings (Author info, clues)
- Optional Data (Rebus, Timer) 

### Header

| Component | Offset | End | Length | Type | Description | 
|:--------------|:-----------|:--------|:-----------|:---------|:----------------|
 | Checksum | 0x00 | 0x01 | 0x2 | short | overall file checksum | 
 | File Magic | 0x02 | 0x0D | 0xC | string | NUL-terminated constant string: 4143 524f 5353 2644 4f57 4e00 ("ACROSS&DOWN") |
| CIB Checksum | 0x0E | 0x0F | 0x2 | short | (defined later) | 
| Masked Low Checksums | 0x10 | 0x13 | 0x4 | | A set of checksums, XOR-masked against a magic string. | 
| Masked High Checksums | 0x14 | 0x17 | 0x4 | | A set of checksums, XOR-masked against a magic string. |
| Version String(?) | 0x18 | 0x1B | 0x4 | string | e.g. "1.2\0" | 
| Reserved1C(?) | 0x1C | 0x1D | 0x2 | ? | In many files, this is uninitialized memory | 
| Scrambled Checksum | 0x1E | 0x1F | 0x2 | short | In scrambled puzzles, a checksum of the real solution (details below). Otherwise, 0x0000. |
| Reserved20(?) | 0x20 | 0x2B | 0xC | ? | In files where Reserved1C is garbage, this is garbage too. | | Width | 0x2C | 0x2C | 0x1 | byte | The width of the board |
 | Height | 0x2D | 0x2D | 0x1 | byte | The height of the board | 
 | # of Clues | 0x2E | 0x2F | 0x2 | short | The number of clues for this board | 
 | Unknown Bitmask | 0x30 | 0x31 | 0x2 | short | A bitmask. Operations unknown. | 
 | Scrambled Tag | 0x32 | 0x33 | 0x2 | short | 0 for unscrambled puzzles. Nonzero (often 4) for scrambled puzzles. |

### Puzzle State

Next come the board solution and player state. (If a player works on a puzzle and then saves their game, the cells they've filled are stored in the state. Otherwise the state is all blank cells and contains a subset of the information in the solution.)

Boards are stored as a single string of ASCII, with one character per cell of the board beginning at the top-left and scanning in reading order, left to right then top to bottom. We'll use this board as a running example (where # represents a black cell, and the letters are the filled-in solution)

``` 
C A T
# # A
# # R
```

At the end of the header (offset 0x34) comes the solution to the puzzle. Non-playable (ie: black) cells are denoted by '.'. So for this example, the board is stored as nine bytes: `CAT..A..R`

Next comes the player state, stored similarly. Empty cells are stored as '-', so the example board before any cells had been filled in is stored as: `---..-..-`


### String Section

Immediately following the boards comes the strings. All strings are encoded in ISO-8859-1 and end with a NUL. Even if a string is empty, its trailing NUL still appears in the file. In order, the strings are:

| Description | Example |
 |:------------|:--------| 
| Title | Theme: .PUZ format |
 | Author | J. Puz / W. Shortz | 
| Copyright | (c) 2007 J. Puz | 
| Clue#1 | Cued, in pool |
 | ... | ...more clues... | 
 | Clue#n | Quiet | 
 | Notes | http://mywebsite |


 ### Clues Assignment

```
#Returns true if the cell at (x, y) gets an "across" clue number.
def cell_needs_across_number(x, y): 
# Check that there is  blank to the left of us 
    if x == 0 or is_black_cell(x-1, y): 
        # Check that there is space (at least two cells) for a word here 
        if x+1 < width and is_black_cell(x+1): 
            return True 
    return False

def cell_needs_down_number(x, y): 
# ...as above, but on the y axis 
```

And then the actual assignment code: 

```
An array mapping across clues to the "clue number".
So across_numbers[2] = 7 means that the 3rd across clue number
points at cell number 7.

across_numbers = []
down_numbers = []
cur_cell_number = 1

# Iterate through the state array
for y in 0..height: 
    for x in 0..width: 
        if is_black_cell(x, y): 
            continue
        assigned_number = False
        if cell_needs_across_number(x, y):
            across_numbers.append(cur_cell_number)
            cell_numbers[x][y] = cur_cell_number
            assigned_number = True
        if cell_needs_down_number(x, y):
            down_numbers.append(cur_cell_number)
            cell_numbers[x][y] = cur_cell_number
            assigned_number = True
        
        if assigned_number:    
        cur_cell_number += 1
```