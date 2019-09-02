package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomocy/desk"
)

func main() {
	cnf := parseConfig()
	if err := desk.Create(cnf.dir, cnf.name); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create desk: %s\n", err)
		os.Exit(1)
	}
}

func parseConfig() config {
	dir := flag.String("dir", "./", "the name of dir for files to be created")
	flag.Parse()

	return config{
		dir: *dir, name: flag.Arg(0),
	}
}

type config struct {
	dir, name string
}
