package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"webhook-runner/webhooked"
)

var srvCommand = cobra.Command{
	Use:   "webhooked",
	Short: "webhooked [-f config.hcl] [-l :3000]",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv, err := webhooked.New()
		if err != nil {
			return err
		}

		if err := srv.Listen(cmd.Flags().Lookup("listen").Value.String()); err != nil {
			return err
		}

		return nil
	},
}

func main() {
	srvCommand.Flags().StringP("config", "c", "config.hcl", "The configuration file to read in")
	srvCommand.Flags().StringP("listen", "l", ":3000", "Which port to listen as")
	srvCommand.Flags().StringP("gopath", "g", "./usr", "The GOPATH with the scripts to load")
	if err := viper.BindPFlags(srvCommand.Flags()); err != nil {
		fmt.Println(err)
		return
	}

	viper.AutomaticEnv()

	if err := srvCommand.Execute(); err != nil {
		os.Exit(1)
		return
	}
}
