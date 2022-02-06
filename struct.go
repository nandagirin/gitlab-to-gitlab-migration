package main

import (
	"time"
)

type ExportStatus struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	ExportStatus      string    `json:"export_status"`
	Links             struct {
		APIURL string `json:"api_url"`
		WebURL string `json:"web_url"`
	} `json:"_links"`
}

type ImportStatus struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	ImportStatus      string    `json:"import_status"`
	CorrelationID     string    `json:"correlation_id"`
	FailedRelations   []struct {
		ID               int         `json:"id"`
		CreatedAt        time.Time   `json:"created_at"`
		ExceptionClass   string      `json:"exception_class"`
		Source           string      `json:"source"`
		ExceptionMessage interface{} `json:"exception_message"`
		RelationName     string      `json:"relation_name"`
	} `json:"failed_relations"`
	ImportError interface{} `json:"import_error"`
}

type Project struct {
	ID                int           `json:"id"`
	Description       string        `json:"description"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	CreatedAt         time.Time     `json:"created_at"`
	DefaultBranch     string        `json:"default_branch"`
	TagList           []interface{} `json:"tag_list"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string        `json:"http_url_to_repo"`
	WebURL            string        `json:"web_url"`
	ReadmeURL         string        `json:"readme_url"`
	AvatarURL         interface{}   `json:"avatar_url"`
	ForksCount        int           `json:"forks_count"`
	StarCount         int           `json:"star_count"`
	LastActivityAt    time.Time     `json:"last_activity_at"`
	Namespace         struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		Path      string      `json:"path"`
		Kind      string      `json:"kind"`
		FullPath  string      `json:"full_path"`
		ParentID  interface{} `json:"parent_id"`
		AvatarURL interface{} `json:"avatar_url"`
		WebURL    string      `json:"web_url"`
	} `json:"namespace"`
	Links struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	EmptyRepo                      bool   `json:"empty_repo"`
	Archived                       bool   `json:"archived"`
	Visibility                     string `json:"visibility"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled       bool   `json:"container_registry_enabled"`
	ContainerExpirationPolicy      struct {
		Cadence       string      `json:"cadence"`
		Enabled       bool        `json:"enabled"`
		KeepN         int         `json:"keep_n"`
		OlderThan     string      `json:"older_than"`
		NameRegex     interface{} `json:"name_regex"`
		NameRegexKeep interface{} `json:"name_regex_keep"`
		NextRunAt     time.Time   `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled              bool        `json:"issues_enabled"`
	MergeRequestsEnabled       bool        `json:"merge_requests_enabled"`
	WikiEnabled                bool        `json:"wiki_enabled"`
	JobsEnabled                bool        `json:"jobs_enabled"`
	SnippetsEnabled            bool        `json:"snippets_enabled"`
	ServiceDeskEnabled         bool        `json:"service_desk_enabled"`
	ServiceDeskAddress         interface{} `json:"service_desk_address"`
	CanCreateMergeRequestIn    bool        `json:"can_create_merge_request_in"`
	IssuesAccessLevel          string      `json:"issues_access_level"`
	RepositoryAccessLevel      string      `json:"repository_access_level"`
	MergeRequestsAccessLevel   string      `json:"merge_requests_access_level"`
	ForkingAccessLevel         string      `json:"forking_access_level"`
	WikiAccessLevel            string      `json:"wiki_access_level"`
	BuildsAccessLevel          string      `json:"builds_access_level"`
	SnippetsAccessLevel        string      `json:"snippets_access_level"`
	PagesAccessLevel           string      `json:"pages_access_level"`
	EmailsDisabled             interface{} `json:"emails_disabled"`
	SharedRunnersEnabled       bool        `json:"shared_runners_enabled"`
	LfsEnabled                 bool        `json:"lfs_enabled"`
	CreatorID                  int         `json:"creator_id"`
	ImportStatus               string      `json:"import_status"`
	OpenIssuesCount            int         `json:"open_issues_count"`
	CiDefaultGitDepth          int         `json:"ci_default_git_depth"`
	PublicJobs                 bool        `json:"public_jobs"`
	BuildTimeout               int         `json:"build_timeout"`
	AutoCancelPendingPipelines string      `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex         interface{} `json:"build_coverage_regex"`
	CiConfigPath               string      `json:"ci_config_path"`
	SharedWithGroups           []struct {
		GroupID          int         `json:"group_id"`
		GroupName        string      `json:"group_name"`
		GroupFullPath    string      `json:"group_full_path"`
		GroupAccessLevel int         `json:"group_access_level"`
		ExpiresAt        interface{} `json:"expires_at"`
	} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool   `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool   `json:"allow_merge_on_skipped_pipeline"`
	RequestAccessEnabled                      bool   `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool   `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool   `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool   `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string `json:"merge_method"`
	SuggestionCommitMessage                   string `json:"suggestion_commit_message"`
	AutoDevopsEnabled                         bool   `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool   `json:"autoclose_referenced_issues"`
	RepositoryStorage                         string `json:"repository_storage"`
}
