package manager

type Record struct {
    Name   string   `yaml:"name"`
    Type   string   `yaml:"type"`
    Values []string `yaml:"values"`
}
