package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IAmRiteshKoushik/attendash/pages"
	"github.com/IAmRiteshKoushik/attendash/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	EndpointUrl string `mapstructure:"ENDPOINT_URL"`
	ProjectKey  string `mapstructure:"PROJECT_KEY"`
	ApiKey      string `mapstructure:"API_KEY"`
	Mode        string `mapstructure:"MODE"`
}

var (
	cfgFile     = ".env.toml"
	userLicense string

	// Load from environment
	cfg *Config

	// Platform specific stuff
	appwriteClient client.Client
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "attendash",
	Short: "Admin dashboard TUI for Attendex (https://github.com/IAmRiteshKoushik/attendex)",
	Long: `Attendash is a terminal-based admin dashboard for managing attendance 
tracker data. It provides a streamlined TUI to view, edit, and analyze attendance 
records for ACM events.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pages.DashboardInit()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initClient, loadLicense)
}

func validateEnv(cfg *Config) error {
	return validation.ValidateStruct(
		cfg,
		validation.Field(&cfg.EndpointUrl, validation.Required, is.URL),
		validation.Field(&cfg.ProjectKey, validation.Required),
		validation.Field(&cfg.ApiKey, validation.Required),
		validation.Field(
			&cfg.Mode,
			validation.Required,
			validation.In("DEV", "PROD"),
		),
	)
}

func initConfig() {
	v := viper.New()

	v.SetConfigFile(cfgFile)
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		panic(utils.ErrorString(fmt.Sprintf("Failed to read config: %v", err)))
	}

	if err := v.Unmarshal(&cfg); err != nil {
		panic(
			utils.ErrorString(
				fmt.Sprintf("Failed to serialize config: %v", err),
			),
		)
	}

	if err := validateEnv(cfg); err != nil {
		panic(utils.ErrorString(fmt.Sprintf("Invalid configuration: %v", err)))
	}
}

func loadLicense() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(
			utils.ErrorString(
				fmt.Sprintf("Could not get working directory: %v", err),
			),
		)
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
		log.Fatal(
			utils.ErrorString(
				fmt.Sprintf(
					"LICENSE file not found or failed to read: %v",
					err,
				),
			),
		)
	}

	userLicense = string(content)
}

func initClient() {
	appwriteClient = client.New(
		appwrite.WithProject(cfg.ProjectKey),
		appwrite.WithKey(cfg.ApiKey),
		appwrite.WithEndpoint(cfg.EndpointUrl),
	)

	Orm = appwrite.NewTablesDB(appwriteClient)
}
