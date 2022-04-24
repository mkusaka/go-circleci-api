/*
Copyright © 2022 mkusaka

*/
package cmd

import (
	"os"

	"github.com/mkusaka/go-circleci-api/cmd/context"
	"github.com/mkusaka/go-circleci-api/cmd/insights"
	"github.com/mkusaka/go-circleci-api/cmd/job"
	"github.com/mkusaka/go-circleci-api/cmd/pipeline"
	"github.com/mkusaka/go-circleci-api/cmd/project"
	"github.com/mkusaka/go-circleci-api/cmd/schedule"
	"github.com/mkusaka/go-circleci-api/cmd/user"
	"github.com/mkusaka/go-circleci-api/cmd/workflow"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-circleci-api",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(context.Cmd)
	rootCmd.AddCommand(insights.Cmd)
	rootCmd.AddCommand(job.Cmd)
	rootCmd.AddCommand(pipeline.Cmd)
	rootCmd.AddCommand(project.Cmd)
	rootCmd.AddCommand(schedule.Cmd)
	rootCmd.AddCommand(user.Cmd)
	rootCmd.AddCommand(workflow.Cmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-circleci-api.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
