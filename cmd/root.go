package main

import (
	"github.com/edatts/banner"
	"github.com/spf13/cobra"
)

const defaultText = "Hello World"

var rootCmd = &cobra.Command{
	Use:          "banner",
	SilenceUsage: true,
	RunE:         printBanner,
}

var inputText string

func init() {
	rootCmd.Flags().StringVarP(&inputText, "text", "t", defaultText, "The input text to create a banner from.")
}

func printBanner(_ *cobra.Command, _ []string) error {
	b, err := banner.NewFromText(inputText)
	if err != nil {
		return err
	}

	b.Print()
	return nil
}
