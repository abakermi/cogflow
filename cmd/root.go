package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"flag"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cogflow",
		Short: "A CLI tool for managing Amazon Cognito",
		Long:  "cogflow is a command-line tool to manage users, groups, and more in Amazon Cognito.",
		Run: func(cmd *cobra.Command, args []string) {
			// Display help by default
			cmd.Help()
		},
	}

	awsProfile string
	region string
	clientID   string
	poolID   string
)

func init() {
	// rootCmd.PersistentFlags().StringVarP(&awsProfile, "profile", "p", "", "AWS profile to use")
	// rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "AWS profile to use")
	// rootCmd.PersistentFlags().StringVarP(&poolID, "pool-id", "pool", "", "Cognito Pool ID")
	// rootCmd.PersistentFlags().StringVarP(&clientID, "client-id", "c", "", "Cognito Client ID")
	flag.Parse()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
