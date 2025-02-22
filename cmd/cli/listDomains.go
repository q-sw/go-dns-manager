package cli

import (
    "github.com/q-sw/go-dns-manager/internal/gandi"

    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// getDevStatusCmd represents the getDevStatus command
var listDomain = &cobra.Command{
    Use:   "domains",
    Short: "Get all domains from Gandi",
    Run: func(cmd *cobra.Command, args []string) {
        d := gandi.ListDomains(viper.GetString("apiURL"), viper.GetString("apiToken"))
        out, _ := json.MarshalIndent(d, "", "    ")
        fmt.Println(string(out))
    },
}

func init() {
    listCmd.AddCommand(listDomain)
}
