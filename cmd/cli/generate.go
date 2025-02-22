package cli

import (
    "github.com/spf13/cobra"
)

// genCmd represents the generate command
var genCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate blank file",
}

func init() {
    rootCmd.AddCommand(genCmd)
}
