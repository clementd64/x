package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

func run() (err error) {
	in := flag.String("i", "", "input file")
	out := flag.String("o", "", "output file")
	flag.Parse()

	if *in == "" || *out == "" {
		flag.Usage()
		return nil
	}

	inFile, err := os.ReadFile(*in)
	if err != nil {
		return err
	}

	var v2 openapi2.T
	if err = json.Unmarshal(inFile, &v2); err != nil {
		return err
	}

	v3, err := openapi2conv.ToV3(&v2)
	if err != nil {
		return err
	}

	outFile, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return json.NewEncoder(outFile).Encode(v3)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
