package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Display the license information for this application",
	Long: `The license command outputs the full licensing terms and conditions 
under which this application is distributed. It helps users and contributors 
understand the legal usage, distribution, and modification rights associated 
with the software, ensuring compliance and transparency.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(userLicense)
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}
