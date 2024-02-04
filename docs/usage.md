## Usage

There are several ways to use Service Level Helper. You can calculate the maximum allowable downtime for a given service level (Error Budget), you can calculate the necessary frequency of probes for monitoring and alerting to achieve a specified service level or you can calculate your service level based on the total duration of outages and the total duration of the monitoring period.

## Help

To get help with the `slh` command, use the `--help` flag or the `help` subcommand.

```bash
$ slh --help
$ slh help
```

## Standard

To use Service Level Helper, pass a service level objective or agreement to the `slh` command. The tool will then calculate the maximum allowable downtime for that period (Error Budget).

Here's an example:

```bash
$ slh 99.9
```

This command will output:

```bash
With a 99.9% Service Level, the maximum downtime allowed is:
Daily (24 hours): 1m 26s
Weekly (7 days): 10m 4s
Monthly (Average): 43m 50s
Quartely (90 days): 2h 9m 35s
Yearly (365 days): 8h 45m 35s
```

## Simple 

You can also use the `slh` command with the number of nines and the `slh` will calculate the maximum allowable downtime for that period (Error Budget).

Here's an example:

```bash
$ slh 4
```

This command will output:

```bash
With a 4 nines or 99.99% Service Level, the maximum downtime allowed is:
Daily (24 hours): 8s
Weekly (7 days): 1m 
Monthly (Average): 4m 23s
Quartely (90 days): 12m 57s
Yearly (365 days): 52m 33s
```

## Probe Frequency Calculation

The Service Level Helper (slh) includes a feature to calculate the necessary frequency of probes for monitoring and alerting to achieve a specified Service Level. This calculation is based on the Mean Time to Repair (MTTR), the expected number of incidents, and the number of probes required to confirm an incident.

To get more information about the probe frequency calculation, visit the [Probe Frequency Calculation Algorithm](probe.md) page.

Here's an example:

```bash
$ slh --mttr 0h20m --incidents 3 --probes 2 99
```

This command will output:

```bash
With a 99% Service Level, the maximum downtime allowed is:
Daily (24 hours): 14m 24s
Weekly (7 days): 1h 40m 48s
Monthly (Average): 7h 18m 20s
Quartely (90 days): 21h 36m 
Yearly (365 days): 3d 15h 36m 

With a 0h20m MTTR, 3 average incident per time period and 2 failed probe to alert,
Here is the MINIMUM probing frequency is necessary to keep your Service Level inside these time periods: 
Weekly (7 days): 6m 48s
Monthly (Average): 1h 3m 3s
Quartely (90 days): 3h 26m 
Yearly (365 days): 14h 26m
```

## Reverse Calculation

You can also use the `slh` command with the total duration of outages.

Here's an example:

```bash
$ slh reverse 5h45m30s
```

```bash
$ slh r 5h45m30s
```
This command will output:

```bash
Service Level for 5h45m30s of downtime:
Daily (24 hours): 76.00694%
Weekly (7 days): 96.57242%
Monthly (Average): 99.21179%
Quartely (90 days): 99.73341%
Yearly (365 days): 99.93427%
```

## Setting a specific time duration for a day

You can also set a custom day duration for all time period calculations with the flag --hours or --hr.

Here's an example:

```bash
$ slh --hours 12 99
```

This command will output:

```bash
With a 99% Service Level, the maximum downtime allowed is:
Daily (12 hours): 7m 12s
Weekly (7 days): 50m 24s
Monthly (Average): 3h 39m 10s
Quartely (90 days): 10h 48m 
Yearly (365 days): 1d 19h 48m 
```

> Keep in mind that the `--hours` flag will set the day duration for all time period calculations (weeky, monthly, quartely and yearly) and also works with reverse calculation.

