package main

import "flag"

func main() {}

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
