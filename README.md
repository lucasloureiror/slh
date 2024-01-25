# Service Level Helper - slh

Service Level Helper (slh) is a command-line interface (CLI) tool designed to assist Site Reliability Engineers (SRE), DevOps professionals, and similar roles in calculating the maximum allowable downtime based on a given service level objective or agreement. This tool is written in Go.

## Features

- Accepts a service level objective or agreement as input.
- Calculates the maximum allowable downtime for daily, weekly, monthly, quarterly, and yearly periods.

## Installation

To install Service Level Helper, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/).

Once you have Go installed, you can install Service Level Helper by running the following command:

```bash
go get github.com/lucasloureiror/slh
```


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

## License

Service Level Helper (slh) is licensed under the [MIT License](LICENSE.md).

