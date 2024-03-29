/*
--------------------------------------------------------------------------------
urfave/cli, dependency used in this project and specifically in this file is made by urfave and is licensed under the MIT License.
--------------------------------------------------------------------------------
Copyright 2024 github.com/lucasloureiror/slh maintainers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"fmt"
	"github.com/lucasloureiror/slh/internal/calculator"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var input calculator.Input

func Start(version string) {

	app := &cli.App{
		Name:                 "slh",
		Usage:                "Service Level Helper is a CLI tool for calculating Service Level related metrics like SLO, SLA, Error Budgets and probing frequency to maintain SLO.",
		Version:              version,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "reverse",
				Aliases: []string{"r"},
				Usage:   "Reverse calculate the Service Level percentage based on total duration of downtime in a string format like 1d or 1h30m or 1h30m15s.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "hours",
						Aliases:     []string{"hr"},
						Usage:       "`number` of hours per day that you want to calculate your availability.",
						Destination: &input.HoursPerDay,
						Value:       24,
						DefaultText: "24",
						Action: func(ctx *cli.Context, h int) error {
							if h < 1 {
								return fmt.Errorf("you need to specify a number of hours greater than 0")
							}

							if h > 24 {
								return fmt.Errorf("you need to specify a number of hours no greater than 24")
							}
							return nil
						},
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 1 {
						return fmt.Errorf("you need to specify the total duration of downtime in the correct format, see help for instructions")
					}
					input.TotalOutageTime = c.Args().First()
					reverseCalculator := calculator.ReverseCalculator{}
					calculator.Start(reverseCalculator, input)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mttr",
				Aliases: []string{"m"},
				Usage: "Use Mean Time to Repair (MTTR) in format (XhYm - see examples below) " +
					"to get a frequency for monitoring and alerting necessary on the service to achieve the Service Level." +
					"\nExample for 30 minutes MTR: slh --mtr 0h30m 99.95" +
					"\nExample for 1 hour MTR: slh --mtr 1h0m 99.95",
				Destination: &input.MTTR,
				Action: func(ctx *cli.Context, m string) error {
					if m == "" {
						return fmt.Errorf("you need to specify a MTTR duration with --mttr or -m")
					}
					return nil
				},
			},
			&cli.IntFlag{
				Name:        "incidents",
				Aliases:     []string{"i"},
				Usage:       "Is the `number` of Incidents that you are expecting to happen in a specific time frame.",
				Destination: &input.Incidents,
				Value:       1,
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
				Usage:       "`number` of probes necessary to alert you about an incident.",
				Destination: &input.ProbeFailures,
				Value:       1,
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
			&cli.IntFlag{
				Name:        "hours",
				Aliases:     []string{"hr"},
				Usage:       "`number` of hours per day that you want to calculate your availability.",
				Destination: &input.HoursPerDay,
				Value:       24,
				DefaultText: "24",
				Action: func(ctx *cli.Context, h int) error {
					if h < 1 {
						return fmt.Errorf("you need to specify a number of hours greater than 0")
					}

					if h > 24 {
						return fmt.Errorf("you need to specify a number of hours no greater than 24")
					}
					return nil
				},
			},
		},

		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 1 {
				return fmt.Errorf("you need to specify one Service Level percentage")
			}
			input.ServiceLevel = ctx.Args().First()
			downtimeCalculator := calculator.DowntimeCalculator{}
			calculator.Start(downtimeCalculator, input)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
