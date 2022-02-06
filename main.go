package main

import (
	"log"
	"os"
	"time"
)

func main() {
	var status bool
	var err error

	// Variable setters
	sourceGitlabUrl := os.Getenv("SOURCE_GITLAB_URL")
	sourceGitlabProjectId := os.Getenv("SOURCE_GITLAB_PROJECT_ID")
	sourceGitlabPrivateToken := os.Getenv("SOURCE_GITLAB_PRIVATE_TOKEN")
	destinationGitlabUrl := os.Getenv("DESTINATION_GITLAB_URL")
	destinationGitlabPrivateToken := os.Getenv("DESTINATION_GITLAB_PRIVATE_TOKEN")
	destinationGitlabNamespaceName := os.Getenv("DESTINATION_GITLAB_NAMESPACE_NAME")
	destinationGitlabProjectName := os.Getenv("DESTINATION_GITLAB_PROJECT_NAME")

	log.Println("############################################")
	log.Println("##        Starting Gitlab Migration       ##")
	log.Println("############################################")
	log.Println("")
	log.Println("")

	sourceGitlabProjectUrl := sourceGitlabUrl + "/api/v4/projects/" + sourceGitlabProjectId
	sourceProjectName, err := GetProjectName(sourceGitlabProjectUrl, sourceGitlabPrivateToken)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Migrating project", sourceProjectName)
	log.Println("")
	log.Println("")

	// Archive source project
	log.Println("###########################################")
	log.Println("# Archive source project")
	log.Println("###########################################")
	if status, err = ProjectArchive(sourceGitlabProjectUrl, sourceGitlabPrivateToken); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} else {
		log.Println("Project archive succeeded")
	}
	log.Println("")
	log.Println("")

	// Export source project
	log.Println("###########################################")
	log.Println("# Export source project")
	log.Println("###########################################")
	if status, err = ProjectExport(sourceGitlabProjectUrl, sourceGitlabPrivateToken); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} else {
		log.Println("Project export is started")
	}
	log.Println("")
	log.Println("")

	// Get export status and wait until export is finished until timeout is reached
	log.Println("###########################################")
	log.Println("# Get source project export status")
	log.Println("###########################################")
	for stay, timeout, exportStatus := true, time.After(300*time.Second), ""; stay; {
		select {
		case <-timeout:
			stay = false
			log.Println("Project export timeout")
			os.Exit(1)
		default:
			if exportStatus, err = GetProjectExportStatus(sourceGitlabProjectUrl, sourceGitlabPrivateToken); err != nil {
				log.Fatalln(err)
				os.Exit(1)
			} else if exportStatus == "none" {
				log.Println("Export status:", exportStatus)
				log.Fatalln("Project export is not started. Exiting script . . .")
				os.Exit(1)
			} else if exportStatus != "none" && exportStatus != "finished" {
				log.Println("Project export is in progress . . .")
				log.Println("Project export status:", exportStatus)
				time.Sleep(3 * time.Second)
			} else {
				log.Println("Project export is finished")
				stay = false
			}
		}
	}
	log.Println("")
	log.Println("")

	// Fetch exported repo as tar.gz file
	log.Println("###########################################")
	log.Println("# Fetch source project export file")
	log.Println("###########################################")
	if status, err = FetchProjectExportFile(sourceGitlabProjectUrl, sourceGitlabPrivateToken); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} else {
		log.Println("Download project export file succeeded")
	}
	log.Println("")
	log.Println("")

	// Import project to destination gitlab repo
	log.Println("###########################################")
	log.Println("# Import project to destination gitlab")
	log.Println("###########################################")
	destinationImportUrl := destinationGitlabUrl + "/api/v4/projects/import"
	status, importedProjectId, err := ProjectImport(destinationImportUrl, destinationGitlabPrivateToken, destinationGitlabNamespaceName, destinationGitlabProjectName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} else {
		log.Println("Project import is started")
	}
	log.Println("")
	log.Println("")

	destinationGitlabProjectUrl := destinationGitlabUrl + "/api/v4/projects/" + importedProjectId

	// Get import status and wait until import is finished until timeout is reached
	log.Println("###########################################")
	log.Println("# Get destination project import status")
	log.Println("###########################################")
	for stay, timeout, importStatus := true, time.After(3600*time.Second), ""; stay; {
		select {
		case <-timeout:
			stay = false
			log.Println("Project import timeout")
			os.Exit(1)
		default:
			if importStatus, err = GetProjectImportStatus(destinationGitlabProjectUrl, destinationGitlabPrivateToken); err != nil {
				log.Fatalln(err)
				os.Exit(1)
			} else if importStatus == "none" || importStatus == "failed" {
				log.Println("Project import status:", importStatus)
				log.Fatalln("Project import is failed or not started")
				os.Exit(1)
			} else if importStatus != "none" && importStatus != "finished" {
				log.Println("Project import is in progress . . .")
				log.Println("Project import status:", importStatus)
				time.Sleep(3 * time.Second)
			} else {
				log.Println("Project import is finished")
				stay = false
			}
		}
	}
	log.Println("")
	log.Println("")

	// Unarchive destination project
	log.Println("###########################################")
	log.Println("# Unarchive destination project")
	log.Println("###########################################")
	if status, err = ProjectUnarchive(destinationGitlabProjectUrl, destinationGitlabPrivateToken); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} else if status == false {
		os.Exit(1)
	} else {
		log.Println("Project unarchive succeeded")
	}
	log.Println("")
	log.Println("")

	log.Println("############################################")
	log.Println("#  Project has been migrated successfully  #")
	log.Println("############################################")
	log.Println("")
	log.Println("")

	destinationGitlabProjectHttpUrl, err := GetProjectHttpUrl(destinationGitlabProjectUrl, destinationGitlabPrivateToken)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Imported project could be accessed in this url", destinationGitlabProjectHttpUrl)
	log.Println("")
}
