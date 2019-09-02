package main

import "os"

func Example() {
	os.Args = []string{"hello", "world"}
	main()
	// Output:
	// +-----------------+
	// |                 |
	// |                 |
	// |                 |
	// |                 |
	// |        S . . E  |
	// |         . . =   |
	// |            . *  |
	// |             =   |
	// |                 |
	// +-----------------+
}
