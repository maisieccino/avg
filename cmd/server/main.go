package main

import (
	"fmt"
	"github.com/mbellgb/avg/pkg/server"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile = ""
	port    = ""
	host    = ""
)

var rootCmd = &cobra.Command{
	Use:   "avg-server",
	Short: "Starts avg web server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(host, port)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.avg-server.yaml)")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "2222", "port to listen on")
	rootCmd.PersistentFlags().StringVarP(&host, "bind-host", "b", "0.0.0.0", "host to bind to")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.SetDefault("port", "2222")
	viper.SetDefault("host", "0.0.0.0")
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
		viper.SetConfigName(".avg-server")
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
