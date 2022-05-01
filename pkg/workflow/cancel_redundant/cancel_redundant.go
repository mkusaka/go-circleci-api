package cancel_redundant

import (
	context2 "context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/mkusaka/go-circleci-api/client"
	"github.com/mkusaka/go-circleci-api/pkg/utils"
)

func RunE(branch string) error {
	ctx := context2.Background()
	userName := utils.CircleProjectUsername

	repoName := utils.CircleProjectReponame

	workflowId := utils.CircleWorkflowId
	token := utils.CircleToken
	targetBranch := branch
	if targetBranch == "" {
		branchName := utils.CircleBranch
		targetBranch = branchName
	}

	opts := &azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{client.CustomPolicy{}},
	}
	c := client.NewCircleciClient(opts)
	workflow, err := c.GetWorkflowByID(ctx, workflowId, nil)
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get workflow failed: %w", err)
	}

	p, err := c.GetPipelineByID(ctx, *workflow.PipelineID, nil)
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: get pipeline failed: %w", err)
	}

	slug, err := utils.FmtProjectSlug(*p.Vcs.ProviderName, userName, repoName)
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: format slug failed: %w", err)
	}
	var copts *client.CircleciClientListPipelinesForProjectOptions
	if targetBranch != "" {
		copts = &client.CircleciClientListPipelinesForProjectOptions{
			Branch: &targetBranch,
		}
	}

	pipes, err := c.ListPipelinesForProject(ctx, slug, copts)
	if err != nil {
		return fmt.Errorf("doCancelRedundantWorkflow: list pipeline failed: %w", err)
	}

	var wg sync.WaitGroup
	for _, item := range pipes.Items {
		state := *item.State
		if state == client.Enum25Created || state == client.Enum25Pending || state == client.Enum25Setup || state == client.Enum25SetupPending {
			wg.Add(1)
			defer wg.Done()

			_, err := c.CancelWorkflow(ctx, *item.ID, nil)
			if err != nil {
				return err
			}

			fmt.Printf("finish cancel workflow :%s\n", *item.ID)
		}
	}

	wg.Wait()

	fmt.Println("finish cancel workflow")
	return nil
}
