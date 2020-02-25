package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "app",
	Version: "1.0.0",
	Short:   "Instagram tools",
	Long:    `Instagram tools`,
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("users", "u", []string{""}, "Instagram users name")
	rootCmd.PersistentFlags().StringSliceP("lists", "l", []string{""}, "Files (list.txt) with list of Instagram users name")
	rootCmd.PersistentFlags().StringSliceP("tags", "t", []string{""}, "Tags")
	rootCmd.PersistentFlags().StringSliceP("webhooks", "w", []string{""}, "Webhooks")
	viper.BindPFlag("users", rootCmd.PersistentFlags().Lookup("users"))
	viper.BindPFlag("lists", rootCmd.PersistentFlags().Lookup("lists"))
	viper.BindPFlag("tags", rootCmd.PersistentFlags().Lookup("tags"))
	viper.BindPFlag("webhooks", rootCmd.PersistentFlags().Lookup("webhooks"))
}

func main() {
	rootCmd.AddCommand(fotoCmd)

	rootCmd.Help()
	rootCmd.Execute()
}
