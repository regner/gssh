package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGcloudArgs(t *testing.T) {
	vm := vmDetails{Zone: "us-east-2", Name: "herp-derp"}
	project := projectDetails{ProjectID: "this-is-project"}

	response := buildGcloudArgs(vm, project)

	expected := []string{
		"beta",
		"compute",
		"ssh",
		"--tunnel-through-iap",
		"--zone",
		"us-east-2",
		"--project",
		"this-is-project",
		"herp-derp",
	}

	assert.Equal(t, expected, response)
}
