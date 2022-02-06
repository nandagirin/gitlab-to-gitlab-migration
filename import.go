package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

func ProjectUnarchive(url string, privateToken string) (bool, error) {
	var project Project

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("POST", url+"/unarchive", headers, nil)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(respBody, &project); err != nil {
		return false, err
	}

	if project.Archived == true {
		log.Println("Archived status:", project.Archived)
		err := errors.New("Project unarchive failed. Archived status should be False")
		return false, err
	}
	log.Println("Archived status:", project.Archived)

	return true, nil
}

func ProjectImport(url string, privateToken string, targetNamespace string, targetProjectName string) (bool, string, error) {
	var importStatus ImportStatus

	form := map[string]string{
		"namespace": targetNamespace,
		"path":      targetProjectName,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for key, val := range form {
		_ = writer.WriteField(key, val)
	}

	file, err := os.Open("./export.tar.gz")
	defer file.Close()
	fileByte, err := writer.CreateFormFile("file", filepath.Base("./export.tar.gz"))
	_, err = io.Copy(fileByte, file)
	if err != nil {
		return false, "", err
	}

	err = writer.Close()
	if err != nil {
		return false, "", err
	}

	if err != nil {
		return false, "", err
	}

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
		"Content-Type":  writer.FormDataContentType(),
	}

	respBody, err := HttpRequest("POST", url, headers, payload)
	if err != nil {
		return false, "", err
	}

	if err := json.Unmarshal(respBody, &importStatus); err != nil {
		return false, "", err
	}

	if importStatus.ImportStatus == "none" || importStatus.ImportStatus == "failed" {
		log.Println("Import status:", importStatus.ImportStatus)
		err := errors.New("Project import failed to start")
		return false, "", err
	}

	return true, strconv.Itoa(importStatus.ID), nil
}

func GetProjectImportStatus(url string, privateToken string) (string, error) {
	var importStatus ImportStatus

	headers := map[string]string{
		"PRIVATE-TOKEN": privateToken,
	}
	respBody, err := HttpRequest("GET", url+"/import", headers, nil)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(respBody, &importStatus); err != nil {
		return "", err
	}

	return importStatus.ImportStatus, nil
}
