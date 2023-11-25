package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tai",
	Short: "A command line tool for managing your task lists",
	Long:  `A command line tool for managing your task lists.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version)
	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(removeCmd)
	// rootCmd.AddCommand(doneCmd)
	// rootCmd.AddCommand(undoneCmd)
	// rootCmd.AddCommand(listCmd)
}
