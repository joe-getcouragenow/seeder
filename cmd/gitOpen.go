package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// gitCatchupCmd represents the catchup command
var gitOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the github web gui at your PR.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git open called")

		// list all the
		result := pkg.Open()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitOpenCmd)
}
