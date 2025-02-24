package cli

import (
	"github.com/q-sw/go-dns-manager/internal/gandi"

	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getDomain = &cobra.Command{
	Use:   "domain",
	Short: "Get all domains from Gandi",
	Run: func(cmd *cobra.Command, args []string) {
		d := gandi.GetDomain(domain, viper.GetString("apiURL"), viper.GetString("apiToken"))
		out, _ := json.MarshalIndent(d, "", "    ")
		fmt.Println(string(out))
	},
}

var domain string

func init() {
	getCmd.AddCommand(getDomain)

	getDomain.Flags().StringVarP(&domain, "name", "n", "", "Get domain information")
}
