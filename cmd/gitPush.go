package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// gitCatchupCmd represents the catchup command
var gitPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push pushes your committed changes to your Origin as a PR.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git push called")

		// list all the
		result := pkg.Push()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitPushCmd)
}
