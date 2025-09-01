package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense = ".attendash.yaml" // Hardcoded for now, to be changed later

	rootCmd = &cobra.Command{
		Use:   "attendash",
		Short: "Admin dashboard TUI for Attendex (https://github.com/IAmRiteshKoushik/attendex)",
		Long: `Attendash is a terminal-based admin dashboard for managing attendance 
tracker data. It provides a streamlined TUI to view, edit, and analyze attendance 
records for ACM events.`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, loadLicense)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to find home directory: %v", err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".attendash")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found or error reading config:", err)
	}
}

func loadLicense() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Could not get working directory: %v", err)
		return
	}
	licensePath := filepath.Join(cwd, "LICENSE")
	content, err := os.ReadFile(licensePath)
	if err != nil {
		home, err2 := os.UserHomeDir()
		if err2 == nil {
			licensePath = filepath.Join(home, "LICENSE")
			content, err = os.ReadFile(licensePath)
		}
	}
	if err != nil {
		log.Printf("LICENSE file not found or failed to read: %v", err)
		return
	}

	userLicense = string(content)
}
