package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/dep"
)

// installCmd represents the install command
var depUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Ininstalls a dep",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dep uninstall called")

		result := pkg.Uninstall()
		fmt.Println(result)
	},
}

func init() {
	depCmd.AddCommand(depUninstallCmd)
}
