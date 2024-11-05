package cmd

import (
	"fmt"
	"log"
	"os"
	filehelpers "sorting-example/helpers/fileHelpers"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "package_statistics [architecture]",
	Short: "package statistics returns the top 10 packages as found in the indices of the debian packages for a specific architecture",
	Long:  "A fast and flexible sorting statistics",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		arch := args[0]
		// The URL containing the content indices
		debianURL := viper.GetString("debianURL")
		// The location of the file to be saved
		filePath := viper.GetString("outputFile")
		fmt.Println("debianURL = ", debianURL)
		fmt.Println("outputFile = ", filePath)
		if err := filehelpers.DownloadFile(debianURL, filePath, arch); err != nil {
			log.Fatalf("Error in Downloading file: %v", err)
		}

		_, err := filehelpers.UnzipAndCreateArrPackages(filePath, arch)
		if err != nil {
			log.Fatalf("Error in unzip file: %v", err)
		}
		duration := time.Since(start)
		log.Printf("Execution Time = %v\n", duration)

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".cobra.yaml", "config file (default is $HOME/.cobra.yaml)")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	fmt.Println(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
