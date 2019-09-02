package drunkenbishop // import "moul.io/drunken-bishop/drunkenbishop"

import (
	"fmt"
	"strings"
)

const (
	RoomWidth   = 17
	RoomHeight  = 9
	Alphabet    = " .o+=*BOX@%&#/^"
	UnknownChar = '!'
	StartY      = 4
	StartX      = 8
	StartCode   = 1000
	EndCode     = 2000
	StartChar   = 'S'
	EndChar     = 'E'

	NW = "00"
	NE = "01"
	SW = "10"
	SE = "11"
)

type pos struct{ Y, X int }
type Room [RoomHeight][RoomWidth]int

func bytesToBitpairs(input []byte) []string {
	bitpairs := []string{}
	for _, byte := range input {
		bin := fmt.Sprintf("%08b", byte)
		bitpairs = append(
			bitpairs,
			bin[6:8],
			bin[4:6],
			bin[2:4],
			bin[0:2],
		)
	}
	return bitpairs
}

func FromBytes(input []byte) Room {
	room := Room{}
	pos := pos{StartY, StartX}
	for _, bitpair := range bytesToBitpairs(input) {
		switch bitpair {
		case NW:
			pos.Y--
			pos.X--
		case NE:
			pos.Y--
			pos.X++
		case SW:
			pos.Y++
			pos.X--
		case SE:
			pos.Y++
			pos.X++
		}
		switch {
		case pos.Y < 0:
			pos.Y = 0
		case pos.Y >= RoomHeight:
			pos.Y = RoomHeight - 1
		}
		switch {
		case pos.X < 0:
			pos.X = 0
		case pos.X >= RoomWidth:
			pos.X = RoomWidth - 1
		}
		room[pos.Y][pos.X]++
	}
	room[StartY][StartX] = StartCode
	room[pos.Y][pos.X] = EndCode
	return room
}

func (r Room) String() string {
	output := "+" + strings.Repeat("-", RoomWidth) + "+\n"
	for _, row := range r {
		line := []byte{}
		for _, col := range row {
			var char byte
			switch {
			case col == StartCode:
				char = StartChar
			case col == EndCode:
				char = EndChar
			case col < len(Alphabet):
				char = Alphabet[col]
			default:
				char = UnknownChar
			}
			line = append(line, char)
		}
		output += "|" + string(line) + "|\n"
	}
	output += "+" + strings.Repeat("-", RoomWidth) + "+"
	return output
}
