package manager

type Record struct {
    Name   string   `yaml:"name"`
    Type   string   `yaml:"type"`
    TTL    int      `yaml:"ttl"`
    Values []string `yaml:"values"`
}
