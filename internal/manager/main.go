package manager

import (
	"fmt"
	"os"

	"github.com/q-sw/go-dns-manager/internal/gandi"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func ApplyRecord(file string, check bool) {
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error to read the file %v \n", file)
	}

	d := yamlDecoder(f)
	apiToken := viper.GetString("apiToken")
	for k, v := range d {
		if gandi.CheckDomainName(k, apiToken) {
			fmt.Println("")
			for i := 0; i < len(v); i++ {
				status, href, changes := gandi.CheckRecord(k, v[i].Name, v[i].Type, v[i].TTL, v[i].Values)
				if status == "na" {
					color.Cyan(fmt.Sprintf("%v.%v No update required", v[i].Name, k))
				} else if status == "update" {
					if !check {
						gandi.UpdateRecord(href, v[i].TTL, v[i].Values)
						fmt.Println("update")
					}
					fmt.Println(href)
					color.Yellow(fmt.Sprintf("%v.%v Need to be update\n", v[i].Name, k))
					color.Yellow(fmt.Sprintf("%v\n", changes))
				} else {
					if !check {
						gandi.CreateRecord(k, v[i].Name, v[i].Type, v[i].TTL, v[i].Values)
						fmt.Println("create")
					}
					color.Green(fmt.Sprintf("%v.%v Need to be create\n", v[i].Name, k))
				}
			}
		}
	}
}

func yamlDecoder(content []byte) map[string][]Record {
	var c map[string][]Record

	err := yaml.Unmarshal(content, &c)
	if err != nil {

		fmt.Println("error to Unmarshal the file")
	}

	return c
}
