package gandi

import (
    "github.com/q-sw/go-dns-manager/internal/utils"
    "encoding/json"
    "fmt"

    "github.com/spf13/viper"
)

func CheckDomainName(domain, apiToken string) bool {
    for _, d := range ListDomains(viper.GetString("apiURL"), apiToken) {
        if d.Fqdn == domain {
            return true
        }
    }
    return false
}

func ListDomains(apiURL, apiToken string) []Domains {
    var h []utils.Header

    h = append(h, utils.Header{Property: "authorization", Value: fmt.Sprintf("Bearer %v", apiToken)})
    b, _ := utils.HttpRequest("GET", fmt.Sprintf("%v/domain/domains", apiURL), []byte{}, h)

    var domains []Domains
    json.Unmarshal(b, &domains)

    return domains
}

func GetDomain(domain, apiURL, apiToken string) Domain {
    var h []utils.Header

    h = append(h, utils.Header{Property: "authorization", Value: fmt.Sprintf("Bearer %v", apiToken)})
    b, _ := utils.HttpRequest("GET", fmt.Sprintf("%v/livedns/domains/%v", apiURL, domain), []byte{}, h)

    fmt.Println(string(b))
    var d Domain
    json.Unmarshal(b, &d)

    return d
}
