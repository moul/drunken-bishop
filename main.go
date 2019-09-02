package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v2"
	"moul.io/drunken-bishop/drunkenbishop"
)

func main() {
	app := cli.App{
		Flags:  []cli.Flag{}, // width, height, bgcolor, fgcolor
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	input := strings.Join(c.Args().Slice(), " ")
	if input == "" {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		input = string(inBytes)
	}
	room := drunkenbishop.FromBytes([]byte(input))
	fmt.Println(room)
	return nil
}
