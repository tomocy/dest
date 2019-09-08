package desk

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	solutionFilename = "main.go"
	testFilename     = "main_test.go"
)

func Create(dir, name string) error {
	adjusted := adjustName(name)
	dir = filepath.Join(dir, adjusted)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	solFname, testFname := filepath.Join(dir, solutionFilename), filepath.Join(dir, testFilename)
	if err := createDeskFiles(solFname, testFname); err != nil {
		return err
	}

	if err := writeSolution(solFname); err != nil {
		return err
	}
	if err := writeTest(testFname); err != nil {
		return err
	}

	return nil
}

func adjustName(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "_")
}

func createDeskFiles(names ...string) error {
	for _, name := range names {
		f, err := os.Create(name)
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

func writeSolution(name string) error {
	return write(name, "package main\n\nfunc main() {}\n\nfunc solve() {}\n")
}

func writeTest(name string) error {
	return write(name, `package main

import "testing"

var solutions = map[string]func(interface{}) interface{}{}

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected interface{}
	}{}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
				if actual != test.expected {
					t.Errorf("got , expect ", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(interface{}{})
			}
		})
	}
}`)
}

func write(name, content string) error {
	dst, err := os.OpenFile(name, os.O_WRONLY, 0700)
	if err != nil {
		return err
	}

	fmt.Fprintf(dst, content)

	return dst.Close()
}

func remove(names ...string) error {
	for _, name := range names {
		if err := os.RemoveAll(name); err != nil {
			return err
		}
	}

	return nil
}
