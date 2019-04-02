package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGcloudvMs(t *testing.T) {
	execCommand = fakeVmsCommand
	defer func() { execCommand = exec.Command }()

	project := projectDetails{
		ProjectID: "irrelevant",
	}

	expectedVMs := []vmDetails{
		vmDetails{
			Zone: "us-central-1",
			Name: "app-central-as72",
		},
	}

	vms, _ := getGcloudVMs(project)

	assert.Equal(t, expectedVMs, vms)

}

func fakeVmsCommand(command string, args ...string) (cmd *exec.Cmd) {
	cs := []string{"-test.run=TestHelperVMs", "--", command}
	cs = append(cs, args...)

	cmd = exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

	return
}

func TestHelperVMs(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	fmt.Fprintf(os.Stdout, "[{\"zone\":\"us-central-1\",\"name\":\"app-central-as72\"}]")

	os.Exit(0)
}

func TestGrtGcloudVMsError(t *testing.T) {
	execCommand = fakeVMsCommandError
	defer func() { execCommand = exec.Command }()

	project := projectDetails{
		ProjectID: "irrelevant",
	}

	_, err := getGcloudVMs(project)

	assert.NotNil(t, err)
}

func fakeVMsCommandError(command string, args ...string) (cmd *exec.Cmd) {
	cs := []string{"-test.run=TestHelperVMsError", "--", command}
	cs = append(cs, args...)

	cmd = exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

	return
}

func TestHelperVMsError(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	fmt.Fprintf(os.Stderr, "This is an error. Handle me!")

	os.Exit(1)
}

func TestFilterVMs(t *testing.T) {
	vms := []vmDetails{
		vmDetails{Zone: "us-central-1", Name: "app-central-as72"},
		vmDetails{Zone: "us-west-1", Name: "db-west-09as"},
		vmDetails{Zone: "eu-west-1", Name: "something-central-a7m2"},
	}

	filtered := filterVMs(vms, "app")

	expectedVMs := []vmDetails{
		vmDetails{Zone: "us-central-1", Name: "app-central-as72"},
	}

	assert.Equal(t, expectedVMs, filtered)

	filtered = filterVMs(vms, "central")

	expectedVMs = []vmDetails{
		vmDetails{Zone: "us-central-1", Name: "app-central-as72"},
		vmDetails{Zone: "eu-west-1", Name: "something-central-a7m2"},
	}

	assert.Equal(t, expectedVMs, filtered)
}

func TestVMDetailsToSurvey(t *testing.T) {
	vms := []vmDetails{
		vmDetails{Zone: "us-central-1", Name: "app-central-as72"},
		vmDetails{Zone: "us-west-1", Name: "db-west-09as"},
		vmDetails{Zone: "eu-west-1", Name: "something-central-a7m2"},
	}

	response := vmDetailsToSurvey(vms)

	expected := []string{
		"(us-central-1) app-central-as72",
		"(us-west-1) db-west-09as",
		"(eu-west-1) something-central-a7m2",
	}

	assert.Equal(t, expected, response)
}

func TestStringToVMDetails(t *testing.T) {
	response := stringToVMDetails("(us-central-1) app-central-as72")

	expected := vmDetails{Zone: "us-central-1", Name: "app-central-as72"}

	assert.Equal(t, expected, response)
}
