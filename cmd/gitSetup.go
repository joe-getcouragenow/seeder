package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// gitCatchupCmd represents the catchup command
var gitSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up your git fork with a remote to Upstream.",
	Long:  "This is needed so that you can Pr to your Origin, and catchup from the Upstream.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git setup called")

		// list all the
		result := pkg.Setup()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitSetupCmd)
}
