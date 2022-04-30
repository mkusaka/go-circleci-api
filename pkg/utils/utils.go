package utils

import (
	"errors"
	"fmt"
	"os"
)

func CircleProjectUsername() (string, error) {
	return orError("CIRCLE_PROJECT_USERNAME")
}

func CircleProjectRepoName() (string, error) {
	return orError("CIRCLE_PROJECT_REPONAME")
}

func CircleBranch() (string, error) {
	return orError("CIRCLE_BRANCH")
}

func CircleUserName() (string, error) {
	return orError("CIRCLE_USERNAME")
}

func CircleWorkflowId() (string, error) {
	return orError("CIRCLE_WORKFLOW_ID")
}

func CircleToken() (string, error) {
	return orError("CIRCLE_TOKEN")
}

func orError(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", errors.New(fmt.Sprintf("orError: required env var %s is not provided", key))
	}
	return val, nil
}
