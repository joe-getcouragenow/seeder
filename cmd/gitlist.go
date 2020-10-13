package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// listCmd represents the list command
var gitListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all git available",
	Long:  " ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git list called")

		// list all the
		result := pkg.List()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitListCmd)
}
