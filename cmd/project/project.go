/*
Copyright © 2022 mkusaka

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd represents the project command
var Cmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

func init() {
	Cmd.AddCommand(GetProjectBySlugCmd)
	Cmd.AddCommand(ListCheckoutKeysCmd)
	Cmd.AddCommand(CreateCheckoutKeyCmd)
	Cmd.AddCommand(GetCheckoutKeyCmd)
	Cmd.AddCommand(DeleteCheckoutKeyCmd)
	Cmd.AddCommand(CreateEnvVarCmd)
	Cmd.AddCommand(ListEnvVarsCmd)
	Cmd.AddCommand(DeleteEnvVarCmd)
	Cmd.AddCommand(GetEnvVarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
