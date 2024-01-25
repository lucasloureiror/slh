package entry

import (
	"github.com/lucasloureiror/slh/internal/calculator"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Start() {

	app := &cli.App{
		Name:  "slh",
		Usage: "Service Level Helper is a CLI tool for calculating service level related data like SLO, SLA and Error Budget",
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() > 0 {
				sla := ctx.Args().First()
				calculator.Start(sla)

			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
