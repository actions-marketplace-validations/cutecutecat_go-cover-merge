package main

import (
	"os"

	"github.com/cutecutecat/go-cover-merge/cov"
)

//goaction:required
//goaction:description Input dir contains coverage files(.out) to be merged
var input_dir = os.Getenv("input_dir")

//goaction:description Output file coverage file(.out) of merged files
//goaction:default Coverage.out
var output_file = os.Getenv("output_file")

func main() {
	cov.MergeCoverage(input_dir, output_file)
}
