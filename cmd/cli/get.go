package cli

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from Gandi",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
