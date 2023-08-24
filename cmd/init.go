package cmd

import (
	internal "github.com/abakermi/cogflow/internal"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Cognito User Pool ID, Client ID, and AWS profile",
	Run: func(cmd *cobra.Command, args []string) {
		// Set up Koanf to read and write configuration
		k := koanf.New(".")

		// Get the user's home directory
		homeDir, err := os.UserHomeDir()
		parser := json.Parser()
		if err != nil {
			internal.LogError("Error getting home directory:", err)
			return
		}
		
		// Configuration file paths and names
		configPath := filepath.Join(homeDir, ".cogflow")
		configFile := "config.json"
		configFilePath := filepath.Join(configPath, configFile)

		// Prompt user for input
		prompt := promptui.Prompt{
			Label: "Cognito User Pool ID:",
		}
		poolID, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		prompt.Label = "Cognito Client ID:"
		clientID, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		prompt.Label = "AWS Profile"
		awsProfile, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		prompt.Label = "AWS Region"
		awsRegion, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Set configuration values in Koanf
		k.Set("pool-id", poolID)
		k.Set("client-id", clientID)
		
		// Set the profile only if provided
		if awsProfile != "" {
			k.Set("aws-profile", awsProfile)
		}

		// Set the region only if provided
		if awsRegion != "" {
			k.Set("region", awsRegion)
		}

		// Marshal Koanf to JSON
		configData, err := k.Marshal(parser)
		if err != nil {
			fmt.Println("Error marshaling config:", err)
			return
		}

		// Write the configuration to the file
		if err := os.WriteFile(configFilePath, configData, os.ModePerm); err != nil {
			fmt.Println("Error writing config:", err)
			return
		}
		
		fmt.Println("User Pool ID, Client ID, and AWS profile initialized.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
