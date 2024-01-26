package entry

import (
	"fmt"
	"github.com/lucasloureiror/slh/internal/calculator"
	"github.com/lucasloureiror/slh/internal/models"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var input models.Input

func Start() {

	app := &cli.App{
		Name:                 "slh",
		Usage:                "Service Level Helper is a CLI tool for calculating service level related data like SLO, SLA and Error Budget",
		Version:              "v0.3.0",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mttr",
				Aliases: []string{"m"},
				Usage: "Use Mean Time to Repair (MTTR) in format (XhYm - see examples below) " +
					"to get a frequency for monitoring and alerting necessary on the service to achieve the Service Level." +
					"\nExample for 30 minutes MTR: slh --mtr 0h30m 99.95" +
					"\nExample for 1 hour MTR: slh --mtr 1h0m 99.95",
				Destination: &input.MTTR,
			},
			&cli.IntFlag{
				Name:        "incidents",
				Aliases:     []string{"i"},
				Usage:       "Is the `number` of Incidents that you are expecting to happen in a specific time frame.",
				Destination: &input.Incidents,
				DefaultText: "1",
				Action: func(ctx *cli.Context, i int) error {
					if i < 1 {
						return fmt.Errorf("you need to specify a number of incidents greater than 0")
					}

					if input.MTTR == "" {
						return fmt.Errorf("you need to specify a MTTR duration with --mttr or -m")
					}
					return nil
				},
			},
			&cli.IntFlag{
				Name:        "probes",
				Aliases:     []string{"p"},
				Usage:       "`Number` of probes necessary to alert you about an incident.",
				Destination: &input.ProbeFailures,
				DefaultText: "1",
				Action: func(ctx *cli.Context, p int) error {
					if p < 1 {
						return fmt.Errorf("you need to specify a number of probes failure to alert greater than 0")
					}

					if input.MTTR == "" {
						return fmt.Errorf("you need to specify a MTTR duration with --mttr or -m")
					}
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 1 {
				return fmt.Errorf("you need to specify only one Service Level percentage")
			}
			input.SLA = ctx.Args().Get(0)
			calculator.Start(&input)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
