package cmd

import (
	cogflow "github.com/abakermi/cogflow/internal"
	"fmt"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage groups in Amazon Cognito",
	Run: func(cmd *cobra.Command, args []string) {
		// This will be executed when the command is run without subcommands
		fmt.Println("Manage groups command")
	},
}

// groupName is used to store the group name flag value
var groupName string

// groupListCmd represents the 'list' subcommand
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available groups",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := cogflow.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := cogflow.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		groups, err := cognitoClient.ListGroups()
		if err != nil {
			fmt.Println("Error listing groups:", err)
			return
		}
		if len(groups) == 0 {
			fmt.Println("No groups found")
			return
		}

		for _, group := range groups {
			fmt.Println(group)
		}
	},
}

// groupCreateCmd represents the 'create' subcommand
var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new group",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := cogflow.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := cogflow.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = cognitoClient.CreateGroup(groupName)
		if err != nil {
			fmt.Println("Error creating group:", err)
			return
		}

		fmt.Printf("Group '%s' created successfully!\n", groupName)
	},
}

// groupDeleteCmd represents the 'delete' subcommand
var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a group",
	Run: func(cmd *cobra.Command, args []string) {
		// Call ValidateConfig to check if cogflow is initialized
		if err := cogflow.ValidateConfig(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		cognitoClient, err := cogflow.NewCognitoClient()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = cognitoClient.DeleteGroup(groupName)
		if err != nil {
			fmt.Println("Error deleting group:", err)
			return
		}

		fmt.Printf("Group '%s' deleted successfully!\n", groupName)
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupListCmd)
	groupCmd.AddCommand(groupCreateCmd)
	groupCmd.AddCommand(groupDeleteCmd)

	groupCreateCmd.Flags().StringVarP(&groupName, "group-name", "g", "", "Group name")
	groupCreateCmd.MarkFlagRequired("group-name")

	groupDeleteCmd.Flags().StringVarP(&groupName, "group-name", "g", "", "Group name")
	groupDeleteCmd.MarkFlagRequired("group-name")
}
