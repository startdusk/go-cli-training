/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// audioCmd represents the audio command
var audioCmd = &cobra.Command{
	Use:   "audio",
	Short: "sets audio as the upload type",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("audio called")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Printf("error retrieving filename: %s\n", err)
			return err
		}
		if filename == "" {
			return errors.New("missing filename")
		}
		fmt.Println("uploading audio file ", filename)
		return nil
	},
}

func init() {
	audioCmd.Flags().StringP("filename", "f", "", "audio file")
	uploadCmd.AddCommand(audioCmd)
}
