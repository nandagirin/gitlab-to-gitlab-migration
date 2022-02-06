package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

func ProjectArchive(url string, privateToken string) (bool, error) {
	var project Project

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("POST", url+"/archive", headers, nil)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(respBody, &project); err != nil {
		return false, err
	}

	if project.Archived != true {
		log.Println("Archived status:", project.Archived)
		err := errors.New("Project archive failed. Archived status should be True")
		return false, err
	}
	log.Println("Archived status:", project.Archived)

	return true, nil
}

func ProjectExport(url string, privateToken string) (bool, error) {
	var exportStatus ExportStatus

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("POST", url+"/export", headers, nil)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(respBody, &exportStatus); err != nil {
		return false, err
	}

	if exportStatus.ExportStatus == "none" {
		log.Println("Export status:", exportStatus.ExportStatus)
		err := errors.New("Project export failed to start")
		return false, err
	}

	return true, nil
}

func GetProjectExportStatus(url string, privateToken string) (string, error) {
	var exportStatus ExportStatus

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("GET", url+"/export", headers, nil)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(respBody, &exportStatus); err != nil {
		return "", err
	}

	return exportStatus.ExportStatus, nil
}

func FetchProjectExportFile(url string, privateToken string) (bool, error) {
	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("GET", url+"/export/download", headers, nil)
	if err != nil {
		return false, err
	}

	// write the whole body at once
	err = ioutil.WriteFile("./export.tar.gz", respBody, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
