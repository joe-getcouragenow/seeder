package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// gitCatchupCmd represents the catchup command
var gitAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add adds all changes and commits them.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git add called")

		// list all the
		result := pkg.Add()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitAddCmd)
}
