package cancel_redundant_workflow

import (
	context2 "context"
	"fmt"

	"github.com/mkusaka/go-circleci-api/client"
	"github.com/mkusaka/go-circleci-api/pkg/utils"
)

func Run(branch string) error {
	ctx := context2.Background()
	orgSlug, err := utils.CircleProjectUsername()
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get project org slug failed: %w", err)
	}
	vcsSlug := "gh"

	repoName, err := utils.CircleProjectRepoName()
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get project reponame failed: %w", err)
	}

	workflowId, err := utils.CircleWorkflowId()
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get workflowId failed: %w", err)
	}

	token, err := utils.CircleToken()
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get token failed: %w", err)
	}

	targetBranch := branch
	if targetBranch == "" {
		branchName, err := utils.CircleBranch()
		if err != nil {
			return fmt.Errorf("doCancelRedundantWorkflow: get branch name failed: %w", err)
		}
		targetBranch = branchName
	}

	c := client.NewAPIClient(client.NewConfiguration())
	currentWorkflow, _, err := c.WorkflowApi.GetWorkflowById(ctx, workflowId).Execute()
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get workflow failed: %w", err)
	}

	pipes, _, err := c.ProjectApi
}
