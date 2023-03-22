/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload [audio|video] [-f|--filename] <filename>",
	Short: "upload an audio or video file",
	Long: `
This command allows you to upload either an audio or video 
file for metadata extraction.
To pass in a filename, use the -f or --filename flag followed 
by the path of the file.

Examples:
	./audiofile-cli upload audio -f audio/beatdoctor.mp3
	./audiofile-cli upload video --filename video/musicvideo.mp4
	`,
}

var (
	Filename = ""
)

func init() {
	uploadCmd.PersistentFlags().StringVarP(&Filename, "filename", "f", "", "file to upload")
	uploadCmd.MarkPersistentFlagRequired("filename")
	rootCmd.AddCommand(uploadCmd)
}
