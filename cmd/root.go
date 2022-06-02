package cmd

import (
	"fmt"
	"gin-boilerplate/pkg/config"
	"gin-boilerplate/pkg/database"
	"gin-boilerplate/pkg/logger"
	"gin-boilerplate/routers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "accountant",
	Short: "An application that helps manage accounts of users",
	Long: `
This is a CLI that enables users to manage their accounts.
You would be able to add credit transactions and debit transactions to various users.
  `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("inside main")
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.SetupConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	db := database.GetDB()
	router := routers.SetupRoute(db)

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

	cobra.CheckErr(rootCmd.Execute())
}
