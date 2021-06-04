package main

import (
	"github.com/meshplus/bitxhub-kit/log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var logger = log.NewWithModule("cmd")

func main() {
	app := cli.NewApp()
	app.Name = "Faucet"
	app.Usage = "Get the erc-20 and bxh node"
	app.Compiled = time.Now()

	// global flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "repo",
			Usage: "Faucet repository path",
		},
	}

	app.Commands = []cli.Command{
		initCMD,
		startCMD,
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}
