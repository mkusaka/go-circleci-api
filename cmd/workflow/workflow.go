/*
Copyright © 2022 mkusaka

*/
package workflow

import (
	"github.com/spf13/cobra"
)

// Cmd represents the workflow command
var Cmd = &cobra.Command{
	Use:   "workflow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	Cmd.AddCommand(GetWorkflowByIdCmd)
	Cmd.AddCommand(ApprovePendingApprovalJobByIdCmd)
	Cmd.AddCommand(CancelWorkflowCmd)
	Cmd.AddCommand(ListWorkflowJobsCmd)
	Cmd.AddCommand(RerunWorkflowCmd)
	Cmd.AddCommand(CancelRedundantWorkflowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
