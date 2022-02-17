package puzgo

//Custom datatype to map the fields in the puz file
//Length or Offset negative means the value is computed or derived
type field struct {
	Length int //in bytes
	Offset int //offset in bytes
}

//Custom datatype to map parts of puz file to logical sections of a puzzle. The string in the map is linked to the
//relevant struct in crossword
type binaryFormat map[string]field

//header contents are stored and mapped here.
var headerFormat = binaryFormat{
	"checksum": field{
		Offset: 0x00,
		Length: 0x2,
	},
	"fileMagic": field{
		Offset: 0x02,
		Length: 0xC,
	},
	"cibChecksum": field{
		Offset: 0x0E,
		Length: 0x2,
	},
	"lowMaskChecksum": field{
		Offset: 0x10,
		Length: 0x4,
	},
	"highMaskChecksum": field{
		Offset: 0x14,
		Length: 0x4,
	},
	"version": field{
		Offset: 0x18,
		Length: 0x4,
	},
	"scrambledChecksum": field{
		Offset: 0x1E,
		Length: 0x2,
	},
	"width": field{
		Offset: 0x2C,
		Length: 0x1,
	},
	"height": field{
		Offset: 0x2D,
		Length: 0x1,
	},
	"noOfClues": field{
		Offset: 0x2E,
		Length: 0x2,
	},
	"scrambledTag": field{
		Offset: 0x32,
		Length: 0x2,
	},
}

var stateFormat = binaryFormat{
	"answerString": field{
		Offset: 0x34,
		Length: -1,
	},
	"playerStateString": field{
		Offset: -1,
		Length: -1,
	},
}
var stringsFormat = binaryFormat{
	"strings": field{
		Offset: -1,
		Length: -1,
	},
}
