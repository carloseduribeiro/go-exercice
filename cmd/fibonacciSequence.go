/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/carloseduribeiro/go-exercise/internal/application/usecase"

	"github.com/spf13/cobra"
)

// fibonacciSequenceCmd represents the fibonacciSequence command
var fibonacciSequenceCmd = &cobra.Command{
	Use:   "fibonacciSequence",
	Short: "Calculates and return the nth fibonacci sequence number",
	Long:  `Calculates and return the nth fibonacci sequence number`,
	Run: func(cmd *cobra.Command, args []string) {
		number, err := cmd.Flags().GetInt("number")
		if err != nil {
			panic(err)
		}
		useCase := usecase.NewGetFibonacci()
		result := useCase.Execute(number)
		fmt.Println(result.Sequence)
	},
}

func init() {
	rootCmd.AddCommand(fibonacciSequenceCmd)
	fibonacciSequenceCmd.Flags().IntP("number", "n", -1, "number to calculate")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fibonacciSequenceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fibonacciSequenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
