# Welcome to Service Level Helper Documentation
[![Go Reference](https://pkg.go.dev/badge/github.com/lucasloureiror/slh.svg)](https://pkg.go.dev/github.com/lucasloureiror/slh)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lucasloureiror/slh)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasloureiror/slh)](https://goreportcard.com/report/github.com/lucasloureiror/slh)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/lucasloureiror/slh)
![GitHub](https://img.shields.io/github/license/lucasloureiror/slh)

Service Level Helper (slh) is a command-line interface (CLI) tool, written in Go, designed to assist Site Reliability Engineers (SRE), DevOps professionals, and similar roles in calculating the maximum allowable downtime based on a given service level objective or agreement.


## Features

- Accepts a service level objective or agreement as input.
- Calculates the maximum allowable downtime for daily (with custom hours per day possible), weekly, monthly, quarterly, and yearly periods.
- Calculates the minimum probing frequency necessary to keep your Service Level inside specific time periods.

