package main

import (
	"log"
	"os"

	"github.com/chyroc/release-go-mod/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:   "release-go-mod",
		Action: internal.Command,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
