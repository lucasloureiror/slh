package entry

import (
	"log"
	"os"

	"github.com/lucasloureiror/slh/internal/calculator"
	"github.com/lucasloureiror/slh/internal/models"
	"github.com/urfave/cli/v2"
)

var input models.Input

func Start() {

	app := &cli.App{
		Name:                 "slh",
		Usage:                "Service Level Helper is a CLI tool for calculating service level related data like SLO, SLA and Error Budget",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "mttr",
				Usage: "Use Mean Time to Repair (MTTR) in format (XhYm - see examples below) " +
					"to get a frequency for monitoring and alerting necessary on the service to achieve the Service Level." +
					"\nExample for 30 minutes MTR: slh --mtr 0h30m 99.95" +
					"\nExample for 1 hour MTR: slh --mtr 1h0m 99.95",
				Destination: &input.MTTR,
			},
		},
		Action: func(ctx *cli.Context) error {


			if ctx.NArg() > 0 {
				input.SLA = ctx.Args().Get(0)
				calculator.Start(&input)

			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
