package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"insta"
	"log"
)

var fotoCmd = &cobra.Command{
	Use:              "foto",
	Version:          "1.0.0",
	Short:            "Foto grabber",
	Long:             `Foto grabber from Instagram account`,
	TraverseChildren: true,
	Run:              fotoRun,
}

func fotoRun(cmd *cobra.Command, args []string) {

	users := viper.GetStringSlice("users")
	for _, user := range users {
		log.Println("[I] Start ByUser", user)
		insta.GetPostsByUser(user)
	}

	lists := viper.GetStringSlice("lists")
	for _, list := range lists {
		log.Println("[I] Start ByList", list)
		insta.GetPostsByUserList(list)
	}

	tags := viper.GetStringSlice("tags")
	for _, tag := range tags {
		log.Println("[I] Start ByTag", tag)
		insta.GetPostsByTag(tag)
	}
	insta.WG.Wait()
}
