package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Version:          "1.0.0",
	Short: "Instagram tools",
	Long:  `Instagram tools`,
	//Run:   run,
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("users", "u", []string{""}, "Instagram users name")
	rootCmd.PersistentFlags().StringSliceP("lists", "l", []string{""}, "Files (list.txt) with list of Instagram users name")
	rootCmd.PersistentFlags().StringSliceP("tags", "t", []string{""}, "Tags")
	viper.BindPFlag("users", rootCmd.PersistentFlags().Lookup("users"))
	viper.BindPFlag("lists", rootCmd.PersistentFlags().Lookup("lists"))
	viper.BindPFlag("tags", rootCmd.PersistentFlags().Lookup("tags"))
}

func Execute() error {
	rootCmd.Help()
	return rootCmd.Execute()
}

func main() {
	rootCmd.AddCommand(fotoCmd)
	rootCmd.AddCommand(userCmd)
	Execute()
}

