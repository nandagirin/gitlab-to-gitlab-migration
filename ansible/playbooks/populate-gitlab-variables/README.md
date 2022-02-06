## Playbooks to populate CI Vars
is a set of playbook for us to populate CI vars in target project in gitlab.com.

### How to run ansible playbook to populate CI vars fetched from source GitLab host
#### Provide ansible inventory
You could follow inventory template provided in folder `inventory/ci-vars-populate/examples`. Please use different folder for each template you are going to provide. Below is an example provided in file `local.yaml`. With this playbook you could migrate variables from any source levels to any destination levels. Levels curently supported are admin, group, and project.

```yaml
ci_variables:
  - source_var_level: project
    source_id: 'XXX' # source project id
    source_var_key: 'DB_HOST'
    dest_var_level: group
    dest_id: 'XXXXXXX' # destination group id
  - source_var_level: admin
    source_id: ci
    source_var_key: 'CONTAINER_REGISTRY_TOKEN'
    dest_var_level: project
    dest_id: 'XXXXXXX'
    dest_var_key: 'CONTAINER_REGISTRY_TOKEN'
  - source_var_level: group
    source_id: 'XXX'
    source_var_key: 'GOPRIVATE'
    dest_var_level: group
    dest_id: 'XXXXXXX'
```

Please note that value of `source_var_key` will be used if variable `dest_var_key` is not specified. Also, if the level is set to `admin`, you need to set the id to `ci`.

#### Run the playbook
You need to provide source and destination gitlab personal access token that have access to either admin, group, or repository. Please follow this [documentation](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token) to create your Gitlab personal access token. Then, provide the token in the command below.

`ansible-playbook -i inventory/ci-vars-populate/examples playbooks/populate-gitlab-variables/populate-variables-from-source-to-destination-gitlab.yaml -e "source_gitlab_personal_access_token=<source gitlab personal access token>" -e "dest_gitlab_personal_access_token=<destination gitlab personal access token>" -vvv` 
