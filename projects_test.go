package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGcloudProjects(t *testing.T) {
	execCommand = fakeProjectsCommand
	defer func() { execCommand = exec.Command }()

	expectedProjects := []projectDetails{
		projectDetails{
			ProjectID: "testing",
		},
	}

	projects := getGcloudProjects()

	assert.Equal(t, expectedProjects, projects)
}

func fakeProjectsCommand(command string, args ...string) (cmd *exec.Cmd) {
	cs := []string{"-test.run=TestHelperProjects", "--", command}
	cs = append(cs, args...)

	cmd = exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

	return
}

func TestHelperProjects(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	fmt.Fprintf(os.Stdout, "[{\"projectID\":\"testing\"}]")

	os.Exit(0)
}

func TestFilterProjects(t *testing.T) {
	projects := []projectDetails{
		projectDetails{ProjectID: "project-one"},
		projectDetails{ProjectID: "project-two"},
		projectDetails{ProjectID: "project-three"},
	}

	filtered := filterProjects(projects, "one")

	expectedProjects := []projectDetails{
		projectDetails{ProjectID: "project-one"},
	}

	assert.Equal(t, filtered, expectedProjects)

	filtered = filterProjects(projects, "project")

	assert.Equal(t, filtered, projects)
}

// func TestFilterProjects(t *testing.T) {
// 	projects :=
// }
