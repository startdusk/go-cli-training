/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"calculator/cmd"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
		return
	}
	cmd.Execute()
}
