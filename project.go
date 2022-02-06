package main

import (
	"encoding/json"
)

func GetProject(url string, privateToken string) (Project, error) {
	project := Project{}

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("GET", url, headers, nil)
	if err != nil {
		return Project{}, err
	}

	if err := json.Unmarshal(respBody, &project); err != nil {
		return Project{}, err
	}

	return project, nil
}

func GetProjectName(url string, privateToken string) (string, error) {
	project, err := GetProject(url, privateToken)
	if err != nil {
		return "", err
	}

	return project.Name, nil
}

func GetProjectHttpUrl(url string, privateToken string) (string, error) {
	project, err := GetProject(url, privateToken)
	if err != nil {
		return "", err
	}

	return project.HTTPURLToRepo, nil
}
