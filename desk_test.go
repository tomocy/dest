package desk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testDataDir = filepath.Join(os.Getenv("GOPATH"), "/src/github.com/tomocy/desk/testdata")

func TestCreate(t *testing.T) {
	type file struct {
		name, content string
	}
	type expected struct {
		solution, test file
	}
	tests := []struct {
		dir, name string
		expected  expected
	}{
		{testDataDir, "Apple is red", expected{
			file{filepath.Join(testDataDir, "apple_is_red", "main.go"), "package main\nfunc main() {}\nfunc solove() {}\n"},
			file{filepath.Join(testDataDir, "apple_is_red", "main_test.go"), "package main\nimport \"testing\"\nfunc TestSolve(t *testing.T) {}\n"},
		}},
		{testDataDir, "Banana is yellow", expected{
			file{filepath.Join(testDataDir, "banana_is_yellow", "main.go"), "package main\nfunc main() {}\nfunc solove() {}\n"},
			file{filepath.Join(testDataDir, "banana_is_yellow", "main_test.go"), "package main\nimport \"testing\"\nfunc TestSolve(t *testing.T) {}\n"},
		}},
	}

	assert := func(expected file) error {
		f, err := os.Open(expected.name)
		if err != nil {
			return err
		}
		defer f.Close()

		actual, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		if string(actual) != expected.content {
			return fmt.Errorf("unexpected content of file: got %q, expect %q", string(actual), expected.content)
		}

		return nil
	}

	for _, test := range tests {
		if err := Create(test.dir, test.name); err != nil {
			t.Errorf("unexpected error: got %s, expect nil\n", err)
		}
		if err := assert(test.expected.solution); err != nil {
			t.Errorf("unexpected solution: %s\n", err)
		}
		if err := assert(test.expected.test); err != nil {
			t.Errorf("unexpected test: %s\n", err)
		}

		remove(
			filepath.Dir(test.expected.solution.name), filepath.Dir(test.expected.test.name),
		)
	}
}
