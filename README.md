# go-dns-manager
CLI written in GO to manage DNS

## Description

This tool helps manage DNS as code, using a declarative approach.

With this tool we can manage DNS from:  
- Gandi
- Google Cloud DNS (coming soon)

## CLI

```plaintext
cli tools to manage DNS as code

Usage:
  dns-manager [command]

Available Commands:
    apply       Apply configurations
        records     Create or Update records
            dns-manager apply records [flags]
                -c, --check               Check before apply
                -f, --file string         file to apply
                -l, --file-list strings   list of files to apply, ex: "file1, file2"
                -h, --help                help for records
    completion  Generate the autocompletion script for the specified shell
    generate    Generate blank file
        config      Generate config file for dns-manager
            dns-manager generate config [flags]
    get         Get information from DNS
        domain      Get all domains from DNS
            dns-manager get domain [flags]
                -h, --help          help for domain
                -n, --name string   Get domain information
    help        Help about any command
    list        list information from Gandi
        domains     Get all domains from Gandi
            dns-manager list domains [flags]
        records     Get all records for a specific domain name
            dns-manager list records [flags]
                -h, --help          help for records
                -n, --name string   domain name ex: test.com
```
> NOTE: The apply command is currently designed to be non-destructive.

### Configuration file

The DNS manager requires a `config.yaml` configuration file in `${HOME}/.config/dns-manager` to function correctly.

```shell
touch ${HOME}/.config/dns-manager/config.yaml
```

`config.yaml`

```yaml
apiToken: " " # Gandi API TOKEN
apiURL: "https://api.gandi.net/v5" # Gandi API URL 
```

## Install the command line

```shell
make install
```
