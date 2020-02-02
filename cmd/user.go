package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"insta"
)

var userCmd = &cobra.Command{
	Use:              "user",
	Version:          "0.0.0",
	Short:            "Under constractions! User grabber",
	Long:             `Under constractions! User grabber from Instagram account`,
	TraverseChildren: true,
	Run:              userRun,
}

func userRun(cmd *cobra.Command, args []string) {

	users := viper.GetStringSlice("users")
	for _, user := range users {
		insta.GetPostsByUser(user)
	}

	lists := viper.GetStringSlice("lists")
	for _, list := range lists {
		insta.GetPostsByUserList(list)
	}

}
