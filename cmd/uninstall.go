package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/dep"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a dep",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstall called")

		result := pkg.Uninstall()
		fmt.Println(result)
	},
}

func init() {
	depCmd.AddCommand(uninstallCmd)

}
