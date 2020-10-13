package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	pkg "github.com/getcouragenow/seeder/pkg/dep"
)

// installCmd represents the install command
var depInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a dep",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dep install called")

		result := pkg.Install()
		fmt.Println(result)
	},
}

func init() {
	depCmd.AddCommand(depInstallCmd)
}
