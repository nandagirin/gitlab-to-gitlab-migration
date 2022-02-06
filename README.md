# GitLab to GitLab Migration
This repo contains script to migrate a GitLab project from one GitLab host to another (e.g self hosted GitLab to GitLab SaaS).

The Golang script will run through the migration process, mainly by calling related APIs in source Gitlab project and destination Gitlab project. You could refer to this [documentation](https://docs.gitlab.com/ee/api/project_import_export.html) for detailed explanation regarding the APIs. The migration itself will execute below process:

`archvie source project` -> `export source project` -> `get export status until finished` -> `download exported project` -> `import project to destination gitlab` -> `unarchive imported project`

To run the script, below environment variables are needed:
- SOURCE_GITLAB_URL: source gitlab url starting with protocol without trailing slash e.g https://gitlab.example.com
- SOURCE_GITLAB_PROJECT_ID: project id of source project, could be found on the repo page
- SOURCE_GITLAB_PRIVATE_TOKEN: gitlab personal access token that has read/write access to source project
- DESTINATION_GITLAB_URL: destination gitlab url starting with protocol without trailing slash e.g https://gitlab.com
- DESTINATION_GITLAB_PRIVATE_TOKEN: gitlab personal access token that has read/write access to destination path where project will be imported
- DESTINATION_GITLAB_NAMESPACE_NAME: project namespace or path before the project name e.g backend/golang
- DESTINATION_GITLAB_PROJECT_NAME: imported project name e.g auth-service

## Running the script locally
### Pre-requisite
#### Install Golang
Refer to this [document](https://golang.org/doc/install) to install golang.

### Run the script
Provide environment variables explained above. Then, on the root folder of this repo, run below command:
```sh
go run *.go
```
