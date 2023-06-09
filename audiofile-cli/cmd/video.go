/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "sets video as the upload type",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("video called")
	},
}

func init() {
	videoCmd.Flags().StringP("filename", "f", "", "video file")
	uploadCmd.AddCommand(videoCmd)
}
