package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestNoTabInUsage(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(usage))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "\t") {
			t.Errorf("line %s start with a tabulation character", line)
		}
	}
	if scanner.Err() != nil {
		t.Errorf("failed to read usage")
	}
}
