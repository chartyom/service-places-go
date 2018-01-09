package cmd

import (
	"fmt"
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	RootCmd = &cobra.Command{
		Use:   "notifications",
		Short: "Librerio Notifications Microservice",
		Long:  "",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.config.yaml)")
	// tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile := flag.String("cert_file", "", "The TLS cert file")
	// keyFile := flag.String("key_file", "", "The TLS key file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		viper.AddConfigPath(".")
		viper.SetConfigName(".config")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
