package gandi

import (
    "github.com/q-sw/go-dns-manager/internal/utils"
    "encoding/json"
    "fmt"
    "slices"

    "github.com/spf13/viper"
)

func CheckRecord(domain, record, recordType string, recordTTL int, values []string) (string, string, [][]string) {


    apiURL := viper.GetString("apiURL")
    apiToken := viper.GetString("apiToken")

    for _, r := range GetRecords(domain, apiURL, apiToken) {
        if r.Name == record && r.Type == recordType && r.TTL == recordTTL && slices.Equal(r.Values, values) {
            return "na", r.Href, [][]string{}
        } else if r.Name == record && r.Type == recordType && r.TTL == recordTTL && (r.TTL != recordTTL || !slices.Equal(r.Values, values)) {
            return "update", r.Href, [][]string{r.Values, values}
        }
    }
    return "create", "", [][]string{values}
}

func GetRecords(domain, apiURL, apiToken string) []Record {

    var h []utils.Header
    h = append(h, utils.Header{Property: "authorization", Value: fmt.Sprintf("Bearer %v", apiToken)})
    b, _ := utils.HttpRequest("GET", fmt.Sprintf("%v/livedns/domains/%v/records", apiURL, domain), []byte{}, h)

    var records []Record

    json.Unmarshal(b, &records)
    return records
}

func CreateRecord(domain, recordName, recordType string, recordTTL int, recordValues []string) []byte {
    type recordValue struct {
        Name   string   `json:"rrset_name"`
        Type   string   `json:"rrset_type"`
        TTL    int      `json:"rrset_ttl"`
        Values []string `json:"rrset_values"`
    }

    apiToken := viper.GetString("apiToken")
    apiURL := viper.GetString("apiURL")
    var h []utils.Header
    h = append(h, utils.Header{Property: "authorization", Value: fmt.Sprintf("Bearer %v", apiToken)})
    h = append(h, utils.Header{Property: "content-type", Value: "application/json"})

    bodyContent, err := json.Marshal(recordValue{Values: recordValues, Name: recordName,TTL: recordTTL, Type: recordType})
    if err != nil {
        fmt.Printf("error during marshal record value")
    }

    b, _ := utils.HttpRequest("POST", fmt.Sprintf("%v/livedns/domains/%v/records", apiURL, domain), bodyContent, h)

    return b

}

func UpdateRecord(recordHref string, recordTTL int, recordValues []string) []byte {

    type recordValue struct {
        Values []string `json:"rrset_values"`
        TTL    int   `json:"rrset_ttl"`
    }

    apiToken := viper.GetString("apiToken")
    var h []utils.Header
    h = append(h, utils.Header{Property: "authorization", Value: fmt.Sprintf("Bearer %v", apiToken)})
    h = append(h, utils.Header{Property: "content-type", Value: "application/json"})

    bodyContent, err := json.Marshal(recordValue{Values: recordValues, TTL: recordTTL})
    if err != nil {
        fmt.Printf("error during marshal record value")
    }

    b, _ := utils.HttpRequest("PUT", recordHref, bodyContent, h)
    return b

}
