/*
Copyright © 2022 mkusaka

*/
package workflow

import (
	"fmt"

	"github.com/mkusaka/go-circleci-api/pkg/workflow/cancel_redundant"
	"github.com/spf13/cobra"
)

// CancelRedundantWorkflowCmd represents the context command
var CancelRedundantWorkflowCmd = &cobra.Command{
	Use:   "cancel_redundant_workflow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			return fmt.Errorf("error executte cancel_redundant_workflow: %w", err)
		}
		err = cancel_redundant.RunE(branch)
		if err != nil {
			return fmt.Errorf("error executte cancel_redundant_workflow: %w", err)
		}
		return nil
	},
}

func init() {
	CancelRedundantWorkflowCmd.Flags().String("branch", "", "target branch. defaults to current branch")
}
