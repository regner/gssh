package main

import (
	"encoding/json"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

type projectDetails struct {
	ProjectID string `json:"projectID"`
}

func getGcloudProjects() (projects []projectDetails, err error) {
	cmd := execCommand("gcloud", "projects", "list", "--format=json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	json.Unmarshal([]byte(output), &projects)

	return
}

func filterProjects(projects []projectDetails, projectFilter string) (filteredProjects []projectDetails) {
	for _, p := range projects {
		if strings.Contains(p.ProjectID, projectFilter) {
			filteredProjects = append(filteredProjects, p)
		}
	}

	return
}

func selectProject(projects []projectDetails) (project projectDetails) {
	var answer string

	options := projectDetailsToSurvey(projects)

	prompt := &survey.Select{
		Message:  "Select your project:",
		Options:  options,
		PageSize: pageSize,
	}

	survey.AskOne(prompt, &answer, nil)

	project = stringToProjectDetails(answer)

	return
}

func projectDetailsToSurvey(projects []projectDetails) (simplifiedProjects []string) {
	for _, project := range projects {
		simplifiedProjects = append(simplifiedProjects, project.ProjectID)
	}

	return
}

func stringToProjectDetails(surveyAnswer string) (project projectDetails) {
	project.ProjectID = surveyAnswer

	return
}
