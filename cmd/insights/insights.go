/*
Copyright © 2022 mkusaka

*/
package insights

import (
	"github.com/spf13/cobra"
)

// Cmd represents the insights command
var Cmd = &cobra.Command{
	Use:   "insights",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	Cmd.AddCommand(GetProjectWorkflowsPageDataCmd)
	Cmd.AddCommand(GetWorkflowTimeseriesCmd)
	Cmd.AddCommand(GetJobTimeseriesCmd)
	Cmd.AddCommand(GetOrgSummaryDataCmd)
	Cmd.AddCommand(GetAllInsightsBranchesCmd)
	Cmd.AddCommand(GetFlakyTestsCmd)
	Cmd.AddCommand(GetProjectWorkflowMetricsCmd)
	Cmd.AddCommand(GetProjectWorkflowRunsCmd)
	Cmd.AddCommand(GetProjectWorkflowJobMetricsCmd)
	Cmd.AddCommand(GetProjectJobRunsCmd)
	Cmd.AddCommand(GetWorkflowSummaryCmd)
	Cmd.AddCommand(GetProjectWorkflowTestMetricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
