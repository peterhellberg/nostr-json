package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/peterhellberg/nostr-json/npub"
)

func main() {
	in, err := parseInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	hex, err := npub.Decode(in.Npub)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	encodeIndentedJSON(os.Stdout, value{
		"names": {
			in.Name: hex,
		},
	})
}

type value map[string]map[string]string

type input struct {
	Name string
	Npub string
}

func parseInput() (input, error) {
	var in input

	flag.StringVar(&in.Name, "name", "", "Name of the user")
	flag.StringVar(&in.Npub, "npub", "", "Public key for the user")

	flag.Parse()

	if in.Name == "" || in.Npub == "" {
		flag.Usage()

		return in, fmt.Errorf("missing input")
	}

	return in, nil
}

func encodeIndentedJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	return enc.Encode(v)
}
