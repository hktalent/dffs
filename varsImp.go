package main

import (
	"flag"
	"fmt"
	"os"
)

//var dataShards = flag.Int("data", 4, "Number of shards to split the data into")
//var parShards = flag.Int("par", 2, "Number of parity shards")

//var dataShards = flag.Int("data", 4, "Number of shards to split the data into, must be below 257.")
//var parShards = flag.Int("par", 2, "Number of parity shards")
//var outDir = flag.String("out", "", "Alternative output directory")

var (
	outFile    = flag.String("out", "", "Alternative output path/file")
	dataShards = flag.Int("data", 4, "Number of shards to split the data into, must be below 257.")
	parShards  = flag.Int("par", 2, "Number of parity shards")
	outDir     = flag.String("out", "", "Alternative output directory")
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(2)
	}
}
