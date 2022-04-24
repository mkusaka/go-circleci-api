/*
Copyright © 2022 mkusaka

*/
package context

import (
	"github.com/spf13/cobra"
)

// Cmd represents the context command
var Cmd = &cobra.Command{
	Use:   "context",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	Cmd.AddCommand(ListContextsCmd)
	Cmd.AddCommand(CreateContextCmd)
	Cmd.AddCommand(DeleteContextCmd)
	Cmd.AddCommand(GetContextCmd)
	Cmd.AddCommand(ListEnvironmentVariablesFromContextCmd)
	Cmd.AddCommand(AddEnvironmentVariableToContextCmd)
	Cmd.AddCommand(DeleteEnvironmentVariableFromContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
