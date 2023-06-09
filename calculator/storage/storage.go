package storage

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func SetValue(folatVal float64) error {
	f, err := os.OpenFile(viper.GetString("filename"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, writeErr := f.WriteString(fmt.Sprintf("%f", folatVal))
	return writeErr
}

func GetValue() float64 {
	dat, err := os.ReadFile(viper.GetString("filename"))
	if err != nil {
		fmt.Println("unable to read from storage")
		return 0
	}
	floatVal, err := strconv.ParseFloat(string(dat), 64)
	if err != nil {
		return 0
	}
	return floatVal
}
