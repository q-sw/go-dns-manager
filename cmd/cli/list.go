package cli

import (
    "github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "list information from Gandi",
}

func init() {
    rootCmd.AddCommand(listCmd)
}
