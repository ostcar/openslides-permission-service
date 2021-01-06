package test

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestCases(t *testing.T) {
	f, err := os.Open("../../testcases.yml")
	if err != nil {
		t.Fatalf("Can not open case file: %v", err)
	}
	defer f.Close()

	var cases Cases
	if err := yaml.NewDecoder(f).Decode(&cases); err != nil {
		t.Fatalf("Can not decode case file: %v", err)
	}

	for _, c := range cases.cases {
		t.Run(c.name, func(t *testing.T) {
			c.Test(t)
		})
	}
}
