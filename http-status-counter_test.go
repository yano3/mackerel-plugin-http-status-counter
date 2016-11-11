package main

import (
	"testing"
)

func TestGraphDefinition(t *testing.T) {
	var httpStatusCounter HTTPStatusCounterPlugin

	graphdef := httpStatusCounter.GraphDefinition()
	if len(graphdef) != 3 {
		t.Errorf("GraphDefinition: %d should be 3", len(graphdef))
	}
}
