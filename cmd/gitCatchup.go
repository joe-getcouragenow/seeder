package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/git"
)

// gitCatchupCmd represents the catchup command
var gitCatchupCmd = &cobra.Command{
	Use:   "catchup",
	Short: "Catchups all git Prs from Upstream",
	Long:  " ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git catchup called")

		// list all the
		result := pkg.Catchup()
		fmt.Println(result)
	},
}

func init() {
	gitCmd.AddCommand(gitCatchupCmd)
}
