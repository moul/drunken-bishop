package drunkenbishop

import "fmt"

func ExampleFromBytes() {
	fingerprint := []byte{
		0xc9, 0xd4, 0x40, 0x21, 0xad, 0x16, 0x69, 0xca,
		0x9a, 0xae, 0xd6, 0x37, 0xae, 0xc6, 0x4d, 0x72,
	}
	fmt.Println(FromBytes(fingerprint))
	// Output:
	// +-----------------+
	// |      .++.       |
	// |      +..o       |
	// |   . o o. .      |
	// |    o oo .       |
	// |   o .  S        |
	// |  o. E           |
	// | .o =            |
	// | ..+ +           |
	// |o...+..          |
	// +-----------------+
}
