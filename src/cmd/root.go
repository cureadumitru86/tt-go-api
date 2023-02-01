package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "godemo",
		Short: "Example golang project.",
		Long:  `Example golang project with the purpose to share the standard structure of a golang api project.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is $HOME/.conf.yaml)")
	rootCmd.PersistentFlags().StringP("database_url", "d", "", "Database connection string for MySQL.")
	viper.BindPFlag("database_url", rootCmd.PersistentFlags().Lookup("database_url"))
	viper.SetDefault("database_url", "Riaan Stegmann riaan.stegmann@tributetech.com")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".godemo")
	}
	viper.SetEnvPrefix("godemo")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
