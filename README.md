# Service Level Helper - slh

[![Go Reference](https://pkg.go.dev/badge/github.com/lucasloureiror/slh.svg)](https://pkg.go.dev/github.com/lucasloureiror/slh)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasloureiror/slh)](https://goreportcard.com/report/github.com/lucasloureiror/slh)
[![GitHub](https://img.shields.io/github/license/lucasloureiror/slh)](LICENSE.md)

Service Level Helper (slh) is a command-line interface (CLI) tool designed to assist Site Reliability Engineers (SRE), DevOps professionals, and similar roles in calculating the maximum allowable downtime based on a given service level objective or agreement. This tool is written in Go.


## Features

- Accepts a service level objective or agreement as input.
- Calculates the maximum allowable downtime for daily, weekly, monthly, quarterly, and yearly periods.

## Installation

To install Service Level Helper, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/).

Once you have Go installed, you can install Service Level Helper by running the following command:

```bash
go install github.com/lucasloureiror/slh@latest
```
This will install Service Level Helper in your `$GOPATH/bin` directory. Make sure that this directory is in your `$PATH` environment variable.

Soon I will also be providing binaries packages for x86-64 and arm.

## Usage

To use Service Level Helper, pass a service level objective or agreement to the `slh` command. The tool will then calculate the maximum allowable downtime for that period.

Here's an example:

```bash
$ slh 99.9
```

This command will output:

```bash
99.9% availability represents the following maximum allowable downtime to meet your Service Level:
Daily: 1m 26s
Weekly: 10m 4s
Monthly (Average): 43m 50s
Quarterly (90 days): 2h 9m 35s
Yearly (365 days): 8h 45m 35s
```

## Probe Frequency Calculation Algorithm

The Service Level Helper (slh) includes a feature to calculate the necessary frequency of probes for monitoring and alerting to achieve a specified Service Level Objective (SLO). This calculation is based on the Mean Time to Repair (MTTR), the expected number of incidents, and the number of probes required to confirm an incident. The algorithm is as follows:

### Formula

$\text{Probe Frequency} = \frac{\text{Total Available Buffer Time}}{\text{Expected Incidents } \times \text{ Number of Probes}}$


### Components

- **Expected Incidents**: The total number of times the system is expected to become unavailable during a specified period.

- **Number of Probes**: This is the count of probes or checks needed to confirm that the system is indeed unavailable and to trigger an alert.

- **Total Available Buffer Time**: This is the time remaining in your downtime budget. It is calculated with this formula: 

    $ \text{Total Available Buffer Time} = \text{Downtime Budget} - (\text{Time to Repair} \times \text{Number of Incidents})$



### Formula Explanation

This formula calculates the average time allocated for each probe to detect an incident. It helps in determining the frequency of health checks necessary to identify system unavailability within the constraints of the specified Service Level Agreement (SLA).

### Usage Example

```bash
$ slh --mttr 0h20m --incidents 3 --probes 2 99
```

### Output

```txt
99% availability represents the following maximum allowable downtime to meet your Service Level:
Daily: 14m 24s
Weekly: 1h 40m 48s
Monthly (Average): 7h 18m 20s
Quartely (90 days): 21h 36m
Yearly (365 days): 3d 15h 36m

With a 0h20 MTTR, 3 average incident per time period and 2 failed probe to alert
Here is the MINIMUM probing frequency is necessary to keep your Service Level inside these time periods:
Weekly: 6m 48s
Monthly (Average): 1h 3m 3s
Quartely (90 days): 3h 26m
Yearly (365 days): 14h 26m
```

### Output Explanation

The output above shows that with a 20-minute MTTR, 3 average incidents per time period, and 2 failed probes to alert, the minimum probing frequency necessary to keep your Service Level inside these time periods is as follows:

- You can't mantain your Service Level on a daily basis with 20 minutes MTTR, 3 average incidents per day and 2 failed probe to alert.

-  You need to probe at least every 6 minutes and 48 seconds to keep your Service Level inside the weekly period with 3 average incidents per week and 2 failed probe to alert.

- You need to probe at least every 1 hour and 3 minutes and 3 seconds to keep your Service Level inside the monthly period with 3 average incidents per month and 2 failed probe to alert.

- You need to probe at least every 3 hours and 26 minutes to keep your Service Level inside the quarterly period with 3 average incidents per quarter and 2 failed probe to alert.

- You need to probe at least every 14 hours and 26 minutes to keep your Service Level inside the yearly period with 3 average incidents per year and 2 failed probe to alert.

<p class="callout warning">This formula is experimental and I offer no guarantee that will work as expected for you!</p>

## License  

Service Level Helper (slh) is licensed under the [GPL-3.0](LICENSE.md).

