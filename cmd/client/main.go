package main

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	url     = ""
	cfgFile = ""
)

var rootCmd = &cobra.Command{
	Use:   "avg",
	Short: "avg client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("avg")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.avg.yaml)")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "localhost:2222", "URL to communicate with")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.SetDefault("url", "localhost:2222")
}

func initConfig() {
	// use cfgfile from flag if it is set
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".avg")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading in config: ", err)
			os.Exit(1)
		}
	}

	viper.SetEnvPrefix("AVG")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
