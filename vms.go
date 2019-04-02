package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

type vmDetails struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

func getGcloudVMs(project projectDetails) (vms []vmDetails, err error) {
	projectArg := fmt.Sprintf("--project=%s", project.ProjectID)

	cmd := execCommand("gcloud", "compute", "instances", "list", projectArg, "--format=json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	json.Unmarshal([]byte(output), &vms)

	for i, vm := range vms {
		splits := strings.Split(vm.Zone, "/")
		vm.Zone = splits[len(splits)-1]
		vms[i] = vm
	}

	return
}

func filterVMs(vms []vmDetails, vmFilter string) (filteredVMs []vmDetails) {
	for _, vm := range vms {
		if strings.Contains(vm.Name, vmFilter) {
			filteredVMs = append(filteredVMs, vm)
		}
	}

	return
}

func selectVM(vms []vmDetails) (vm vmDetails) {
	var answer string

	options := vmDetailsToSurvey(vms)

	prompt := &survey.Select{
		Message:  "Select your VM:",
		Options:  options,
		PageSize: pageSize,
	}

	survey.AskOne(prompt, &answer, nil)

	vm = stringToVMDetails(answer)

	return
}

func vmDetailsToSurvey(vms []vmDetails) (simplifiedVMs []string) {
	for _, vm := range vms {
		simplifiedVMs = append(simplifiedVMs, fmt.Sprintf("(%s) %s", vm.Zone, vm.Name))
	}

	return
}

func stringToVMDetails(surveyAnswer string) (vm vmDetails) {
	r, _ := regexp.Compile("\\((.*)\\) (.*)")
	results := r.FindStringSubmatch(surveyAnswer)

	// Quick reminder as this always catches me, the first
	// result (0) is the whole string string, not group match.
	vm.Zone = results[1]
	vm.Name = results[2]

	return
}
