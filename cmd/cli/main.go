package cli

import (
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "path/filepath"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "dns-manager",
    Short: "cli tools to manage DNS as code",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

    home, _ := os.UserHomeDir()
    viper.AddConfigPath(filepath.Join(home, ".config", "dns-manager"))
    viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.SetEnvPrefix("dns")
    viper.AutomaticEnv()

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        viper.ConfigFileUsed()
    }
}
