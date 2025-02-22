package cli

import (
    "github.com/q-sw/go-dns-manager/internal/gandi"

    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var listRecords = &cobra.Command{
    Use:   "records",
    Short: "Get all records for a specific domain name",
    Run: func(cmd *cobra.Command, args []string) {
        records := gandi.GetRecords(domainName, viper.GetString("apiURL"), viper.GetString("apiToken"))

        out, _ := json.MarshalIndent(records, "", "    ")
        fmt.Println(string(out))
    },
}

var domainName string

func init() {
    listCmd.AddCommand(listRecords)

    listRecords.Flags().StringVarP(&domainName, "name", "n", "", "domain name ex: test.com")
}
