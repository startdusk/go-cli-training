/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// piperCmd represents the piper command
var piperCmd = &cobra.Command{
	Use:   "piper",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("internal error:", err)
			return
		}
		fmt.Printf("piped in: %s\n", s)
	},
}

func init() {
	rootCmd.AddCommand(piperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// piperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// piperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
