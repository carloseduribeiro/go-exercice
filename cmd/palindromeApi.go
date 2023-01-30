/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/carloseduribeiro/go-exercise/internal/infra/handlers"
	"net/http"

	"github.com/spf13/cobra"
)

// palindromeApiCmd represents the palindromeApi command
var palindromeApiCmd = &cobra.Command{
	Use:   "palindromeApi",
	Short: "Run an api to check if a word is a palindrome",
	Long:  `Run an api to check if a word is a palindrome. By default, the API runs at port 80.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		addr := fmt.Sprintf(":%s", port)
		fmt.Printf("running on port %s", port)
		err = http.ListenAndServe(addr, handlers.NewValidatePalindrome(http.MethodGet, "/palindrome/validate"))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(palindromeApiCmd)
	palindromeApiCmd.Flags().StringP("port", "p", "80", "port to expose the resources")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// palindromeApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// palindromeApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
