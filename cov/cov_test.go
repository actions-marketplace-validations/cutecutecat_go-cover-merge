package cov

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

func TestMergeFiles(t *testing.T) {
	test_dir := path.Join("..", "test")

	cases, err := ioutil.ReadDir(test_dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, case_name := range cases {
		case_path := path.Join(test_dir, case_name.Name())

		tmpDir, err := os.MkdirTemp(".", "*")
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer os.RemoveAll(tmpDir)

		input_dir := path.Join(case_path, "input")
		output_file := path.Join(tmpDir, "coverage_all.out")
		expect_file := path.Join(case_path, "output", "coverage_sum.out")

		MergeCoverage(input_dir, output_file)

		f1, err := ioutil.ReadFile(output_file)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		f2, err := ioutil.ReadFile(expect_file)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		if !bytes.Equal(f1, f2) {
			t.Error("generate files not equal")
		}
	}
}
