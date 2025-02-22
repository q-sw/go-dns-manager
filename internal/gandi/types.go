package gandi

import (
    "time"
)

type Domains struct {
    Dates struct {
        RegistryCreatedAt time.Time `json:"registry_created_at"`
        RegistryEndsAt    time.Time `json:"registry_ends_at"`
    } `json:"dates"`
    Fqdn       string `json:"fqdn"`
    Owner      string `json:"owner"`
    Nameserver struct {
        Current string `json:"current"`
    } `json:"nameserver"`
}

type Domain struct {
    Fqdn               string `json:"fqdn"`
    DomainHref         string `json:"domain_href"`
    DomainRecordsHref  string `json:"domain_records_href"`
    DomainKeysHref     string `json:"domain_keys_href"`
    AutomaticSnapshots bool   `json:"automatic_snapshots"`
}

type Record struct {
    Name   string   `json:"rrset_name"`
    TTL    int      `json:"rrset_ttl"`
    Type   string   `json:"rrset_type"`
    Values []string `json:"rrset_values"`
    Href   string   `json:"rrset_href"`
}
