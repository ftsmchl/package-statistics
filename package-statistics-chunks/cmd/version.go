package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [version]",
	Short: "Print the version number of package_statistics",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)
		fmt.Println("package_statics v0.1")
	},
}
