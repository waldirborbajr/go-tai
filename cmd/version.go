/*
Go blueprint version
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TaiVersion is the version of the cli to be overwritten by goreleaser in the CI run with the version of the release in github
var TaiVersion string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `The version command provides information about the application's version.
Use this command to check the current version of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("T.A.I CLI version %v\n", TaiVersion)
	},
}
