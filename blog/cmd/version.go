package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the blog app server version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
