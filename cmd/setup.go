/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/epikoder/klass-tool/src/models"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db_.Migrator().DropTable(&models.User{})
		db_.AutoMigrate(&models.User{})
		db_.Create(&models.User{
			Username: "test",
			Password: "password",
		})
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
