package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Syntax: %s <infile> <outfile> [key[=value]] ...\n", os.Args[0])
		os.Exit(1)
	}
	inName := os.Args[1]
	outName := os.Args[2]
	args := os.Args[3:]
	var err error
	var in *os.File
	if in, err = os.Open(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v %s\n", inName, err)
		os.Exit(1)
	}
	defer in.Close()
	var b []byte
	if b, err = ioutil.ReadAll(in); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v %s\n", inName, err)
		os.Exit(1)
	}
	var t *template.Template
	if t, err = template.New("").Parse(string(b)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v %s\n", inName, err)
		os.Exit(1)
	}
	var out *os.File
	if out, err = os.Create(os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v %s\n", outName, err)
		os.Exit(1)
	}
	defer out.Close()
	m := make(map[string]string)
	for _, a := range args {
		if strings.Contains(a, "=") {
			b := strings.SplitN(a, "=", 2)
			m[b[0]] = b[1]
		} else {
			m[a] = "true"
		}
	}
	if err = t.Execute(out, m); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v %s\n", outName, err)
		os.Exit(1)
	}
}
