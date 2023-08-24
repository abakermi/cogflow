package cmd

import (
	internal "github.com/abakermi/cogflow/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users in Amazon Cognito",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Any pre-run setup can be done here
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// This will be executed when the command is run without subcommands
		fmt.Println("Manage users command")
	},
}

// userID is used to store the user ID flag value
var userID string

// userGetCmd represents the 'get' subcommand
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user details",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := internal.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := internal.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		user, err := cognitoClient.GetUserByID(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printUserDetails(user)
	},
}

// userDisableCmd represents the 'disable' subcommand
var userDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := internal.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := internal.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = cognitoClient.DisableUser(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("User disabled successfully.")
	},
}

// userEnableCmd represents the 'enable' subcommand
var userEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := internal.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := internal.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = cognitoClient.EnableUser(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("User enabled successfully.")
	},
}

// userDeleteCmd represents the 'delete' subcommand
var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := internal.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := internal.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = cognitoClient.DeleteUser(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("User deleted successfully.")
	},
}

// Function to print user details
func printUserDetails(user *internal.User) {
	if len(user.UserAttributes) > 0 {
		fmt.Println("User Attributes:")
		for _, attr := range user.UserAttributes {
			fmt.Printf("%s: %s\n", *attr.Name, *attr.Value)
		}
	}
}

func init() {
	rootCmd.AddCommand(userCmd)
	userGetCmd.Flags().StringVarP(&userID, "user-id", "u", "", "User ID")
	userGetCmd.MarkFlagRequired("user-id")

	userDisableCmd.Flags().StringVarP(&userID, "user-id", "u", "", "User ID")
	userDisableCmd.MarkFlagRequired("user-id")

	userEnableCmd.Flags().StringVarP(&userID, "user-id", "u", "", "User ID")
	userEnableCmd.MarkFlagRequired("user-id")

	userDeleteCmd.Flags().StringVarP(&userID, "user-id", "u", "", "User ID")
	userDeleteCmd.MarkFlagRequired("user-id")

	userCmd.AddCommand(userDisableCmd)
	userCmd.AddCommand(userEnableCmd)
	userCmd.AddCommand(userDeleteCmd)
	userCmd.AddCommand(userGetCmd)
}
