package main

import (
	"flag"
	"github.com/foxbot/rudolph/generator"
	"github.com/foxbot/rudolph/server"
	"log"
	"time"
)

const hostFlagUsage = "Address to host the HTTP API on"
const workerFlagUsage = "ID of this worker - must fit within [" + string(minUint14) + "," + string(maxUint14) + "]"
const epochFlagUsage = "Epoch of this generator expressed as RFC3339"
const serverlessFlagUsage = "If true, this generator will output one snowflake to stdout"

const minUint14 = 0
const maxUint14 = 1<<14 - 1

var hostFlag string
var workerFlag int
var epochFlag string
var serverlessFlag bool

func init() {
	flag.StringVar(&hostFlag, "host", ":5000", hostFlagUsage)
	flag.IntVar(&workerFlag, "worker", 0, workerFlagUsage)
	flag.StringVar(&epochFlag, "epoch", time.Now().Format(time.RFC3339), epochFlagUsage)
	flag.BoolVar(&serverlessFlag, "serverless", false, serverlessFlagUsage)
	flag.Parse()
}

func main() {
	epoch, err := time.Parse(time.RFC3339, epochFlag)
	if err != nil {
		log.Fatalln(err)
	}
	if workerFlag < minUint14 && workerFlag > maxUint14 {
		log.Fatalln("worker must fit within [", minUint14, ",", maxUint14, "]")
	}
	worker := uint16(workerFlag)

	gen := generator.NewSnowflakeGenerator(epoch, worker)

	if serverlessFlag {
		id := gen.Generate()
		println(id)
		return
	}

	srv := server.NewServer(hostFlag, gen)
	log.Fatalln(srv.Run())
}
