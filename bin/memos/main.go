package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tuantran0910/memos_clone/server/profile"
	"github.com/tuantran0910/memos_clone/server/version"
)

var rootCmd = &cobra.Command{
	Use:   "memos",
	Short: `An open source, lightweight note-taking service. Easily capture and share your great thoughts.`,
	Run: func(_ *cobra.Command, _ []string) {
		instanceProfile := &profile.Profile{
			Mode:        viper.GetString("mode"),
			Addr:        viper.GetString("addr"),
			Port:        viper.GetInt("port"),
			Data:        viper.GetString("data"),
			Driver:      viper.GetString("driver"),
			DSN:         viper.GetString("dsn"),
			InstanceURL: viper.GetString("instance-url"),
			Version:     version.GetCurrentVersion(viper.GetString("mode")),
		}
		fmt.Println(instanceProfile)
	},
}
