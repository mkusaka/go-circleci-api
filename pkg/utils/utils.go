package utils

import (
	"errors"
	"fmt"
	"os"
)

var CircleBranch = orDefault("CIRCLE_BRANCH", "")
var CircleBuildUrl = orDefault("CIRCLE_BUILD_URL", "")
var CircleJob = orDefault("CIRCLE_JOB", "")
var CircleOidcToken = orDefault("CIRCLE_OIDC_TOKEN", "")
var CirclePrReponame = orDefault("CIRCLE_PR_REPONAME", "")
var CirclePrUsername = orDefault("CIRCLE_PR_USERNAME", "")
var CircleProjectReponame = orDefault("CIRCLE_PROJECT_REPONAME", "")
var CircleProjectUsername = orDefault("CIRCLE_PROJECT_USERNAME", "")
var CirclePullRequest = orDefault("CIRCLE_PULL_REQUEST", "")
var CircleRepositoryUrl = orDefault("CIRCLE_REPOSITORY_URL", "")
var CircleSha1 = orDefault("CIRCLE_SHA1", "")
var CircleTag = orDefault("CIRCLE_TAG", "")
var CircleUsername = orDefault("CIRCLE_USERNAME", "")
var CircleWorkflowId = orDefault("CIRCLE_WORKFLOW_ID", "")
var CircleWorkflowJobId = orDefault("CIRCLE_WORKFLOW_JOB_ID", "")
var CircleWorkflowWorkspaceId = orDefault("CIRCLE_WORKFLOW_WORKSPACE_ID", "")
var CircleWorkingDirectory = orDefault("CIRCLE_WORKING_DIRECTORY", "")
var CircleInternalTaskData = orDefault("CIRCLE_INTERNAL_TASK_DATA", "")
var CircleCompareUrl = orDefault("CIRCLE_COMPARE_URL", "")

var CircleToken = orDefault("CIRCLE_TOKEN", "")

func orDefault(key string, d string) string {
	val := os.Getenv(key)
	if val == "" {
		return d
	}
	return val
}

func FmtProjectSlug(vcsName string, userName string, repoName string) (string, error) {
	if vcsName == "Github" || vcsName == "gh" {
		return fmt.Sprintf("gh/%s/%s", userName, repoName), nil
	}
	return "", errors.New(fmt.Sprintf("utils: unsupported format: %s", vcsName))
}
