package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	pageSize = 12
)

var (
	version string
	commit  string
)

var execCommand = exec.Command

func main() {
	var projectFilter string
	var vmFilter string
	var printVersion bool

	flag.StringVar(&projectFilter, "p", "", "Project name, can be partial match")
	flag.StringVar(&vmFilter, "m", "", "VM name, can be partial match")

	flag.BoolVar(&printVersion, "v", false, "Print the version")

	flag.Parse()

	if printVersion {
		fmt.Printf("gssh, version %s", version)
		return
	}

	var selectedProject projectDetails
	var selectedVM vmDetails

	projects := getGcloudProjects()
	filteredProjects := filterProjects(projects, projectFilter)

	if len(filteredProjects) == 0 {
		fmt.Print("No projects found based on project filter.")
		return
	} else if len(filteredProjects) == 1 {
		selectedProject = filteredProjects[0]
	} else {
		selectedProject = selectProject(filteredProjects)
	}

	vms := getGcloudVMs(selectedProject)
	filteredVMs := filterVMs(vms, vmFilter)

	if len(filteredVMs) == 0 {
		fmt.Print("No VMs found based on project and VM filter.")
		return
	} else if len(filteredVMs) == 1 {
		selectedVM = filteredVMs[0]
	} else {
		selectedVM = selectVM(filteredVMs)
	}

	fmt.Printf("Using %s as selected project.\n", selectedProject)
	fmt.Printf("Using %s as selected VM in zone %s.\n", selectedVM.Name, selectedVM.Zone)

	args := buildGcloudArgs(selectedVM, selectedProject)

	cmd := exec.Command("gcloud", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()

	// TODO: Add tests
}

func buildGcloudArgs(selectedVM vmDetails, selectedProject projectDetails) (args []string) {
	args = []string{
		"beta",
		"compute",
		"ssh",
		"--tunnel-through-iap",
		"--zone",
		selectedVM.Zone,
		"--project",
		selectedProject.ProjectID,
		selectedVM.Name,
	}

	return
}
