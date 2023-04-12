/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apiserverCmd represents the apiserver command
var apiserverCmd = &cobra.Command{
	Use:   "apiserver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is graphql_generated.go CLI library for Go that empowers applications.
This application is graphql_generated.go tool to generate the needed files
to quickly create graphql_generated.go Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apiserver called")
	},
}

func init() {
	rootCmd.AddCommand(apiserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}