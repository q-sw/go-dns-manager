package cli

import (
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply configurations",
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
