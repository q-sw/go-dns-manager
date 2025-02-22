package cli

import (
    "fmt"
    "os"

    "github.com/q-sw/go-dns-manager/internal/manager"
    "github.com/spf13/cobra"
)

var applyRecords = &cobra.Command{
    Use:   "records",
    Short: "Create or Update records",
    Run: func(cmd *cobra.Command, args []string) {
        check, _ := cmd.Flags().GetBool("check")
        if file != "" {
            fmt.Println("Apply records")
            if _, err := os.Stat(file); err != nil {
                fmt.Println("Check the file path")
            }
            manager.ApplyRecord(file, check)

        } else {
            for _, f := range files {
                if _, err := os.Stat(f); err != nil {
                    fmt.Printf("Check the file path for %v\n", f)
                }
                manager.ApplyRecord(f, check)

            }
        }
    },
}

var file string
var files []string

func init() {
    applyCmd.AddCommand(applyRecords)

    applyRecords.Flags().BoolP("check", "c", false, "Check before apply")
    applyRecords.Flags().StringVarP(&file, "file", "f", "", "file to apply")
    applyRecords.Flags().StringSliceVarP(&files, "file-list", "l", []string{}, "list of files to apply, ex: \"file1, file2\"")
    applyRecords.MarkFlagsOneRequired("file", "file-list")
    applyRecords.MarkFlagsMutuallyExclusive("file", "file-list")
}
