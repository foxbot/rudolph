# rudolph

Rudolph is a snowflake generator.

### Installation

Rudolph requires:
- go1.8

To install:
1. `go get github.com/foxbot/rudolph`
2. `go install github.com/rudolph`

### Usage

#### Starting

Full command-line usage can be invoked with `rudolph --help`

`-host`: The host address Rudolph's HTTP server will be ran on
`-serverless`: If true, Rudolph will not spawn an HTTP server, and a Snowflake will be written to stdout
`-epoch`: The RFC3339 formatted Epoch for the Snowflake Generator
`-worker`: The Worker ID of this instance (e.g. if running in clusters)

#### Accessing

Rudolph's HTTP API exposes a single endpoint, `GET /snowflake`

This endpoint will return the generated snowflake, as a single string.
JSON was not used to prevent overhead on both the server and consumer side.