package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/dep"
)

// listCmd represents the list command
var depListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all deps available",
	Long:  " ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dep list called")

		// list all the
		result := pkg.List()
		fmt.Println(result)
	},
}

func init() {
	depCmd.AddCommand(depListCmd)
}
