package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var (
	k      = koanf.New(".")
	parser = json.Parser()
)

func init() {
	// Set up Viper to read from $HOME/.cogflow/config.json
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	configPath := filepath.Join(homeDir, ".cogflow")
	configFile := "config.json"
	configFilePath := filepath.Join(configPath, configFile)

	k.Load(file.Provider(configFilePath), parser)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// Customize other logger settings here if needed
}

// LogError logs an error with the logger.
func LogError(message string, err error) {
	logger.WithError(err).Error(message)
}

// LogInfo logs an informational message with the logger.
func LogInfo(message string) {
	logger.Info(message)
}

func ValidateRegion(region string) error {
	// Implement region validation logic here
	// Return an error if the region is not valid
	if region != "us-east-1" && region != "us-west-2" {
		return fmt.Errorf("invalid region")
	}
	return nil
}

// ValidateConfig checks if the config has been initialized.
func ValidateConfig() error {
	if k.String("pool-id") == "" || k.String("client-id") == "" {
		return fmt.Errorf("cogflow not initialized")
	}
	return nil
}
func GetPoolID() string {
	return k.String("pool-id")
}

func GetClientID() string {
	return k.String("client-id")
}

func GetAWSProfile() string {
	return k.String("aws-profile")
}
func GetRegion() string {
	return k.String("region")
}
