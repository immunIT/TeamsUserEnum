package cmd

import (
	"errors"

	"TeamsUserEnum/src/teams"

	"github.com/spf13/cobra"
)

var emailFile string
var email string
var token string

// userenumCmd represents the userenum command
var userenumCmd = &cobra.Command{
	Use:   "userenum",
	Short: "User enumeration on Microsoft Teams",
	Long: `Users can be enumerated on Microsoft Teams with the search features.
This tool validates an email address or a list of email addresses.
If these emails exist the presence of the user is retrieved as well as the device used to connect`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if emailFile == "" && email == "" {
			return errors.New("Argument -f or -e required")
		} else if emailFile != "" && email != "" {
			return errors.New("Only argument -f or -e should be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token = "Bearer " + token
		if email != "" {
			teams.Enumuser(email, token, verbose)
		} else {
			teams.Parsefile(emailFile, token, verbose)
		}
	},
}

func init() {
	rootCmd.AddCommand(userenumCmd)

	userenumCmd.Flags().StringVarP(&emailFile, "file", "f", "", "File containing the email address")
	userenumCmd.Flags().StringVarP(&email, "email", "e", "", "Email address")
	userenumCmd.Flags().StringVarP(&token, "token", "t", "", "Bearer token (only the base64 part: eyJ0...)")
	userenumCmd.MarkFlagRequired("token")
}
