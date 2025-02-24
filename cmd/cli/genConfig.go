package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genConfig = &cobra.Command{
	Use:   "config",
	Short: "Generate config file for dns-manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate config file")
	},
}

func init() {
	genCmd.AddCommand(genConfig)
}
