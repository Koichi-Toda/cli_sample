/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// variable for binding flag 'wait'
var wait int

// dummy job to run while outputting progress bar
func dummyJob(wait int) {
	fmt.Printf("[calling dummy job form %v %v, wait=%v]\n", firstName, lastName, wait)

	count := 1000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
	bar.Finish()
	fmt.Printf("finished!\n")
}

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "calling dummy job",
	Long:  "calling dummy job for first development",
	RunE: func(cmd *cobra.Command, args []string) error {

		// output ascii art to show environment
		myFigure := figure.NewFigure(config.Environment, "", true)
		myFigure.Print()

		var msg string
		switch config.Language {
		case "english":
			msg = "Hello!"
		case "japanese":
			msg = "こんにちは"
		}

		// output hello message based on language specification
		fmt.Printf("%v %v!\n", msg, firstName)

		dummyJob(wait)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(callCmd)

	callCmd.Flags().IntVarP(&wait, "wait", "w", 0, "wait for job interval")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
