/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	nomad "github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

type listFlags struct {
	manifestPath string
}

var lFlags listFlags

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list modules",
	Long:  `list the modules that is managed by MarcoCD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		nc, err := nomad.NewClient(nomad.DefaultConfig())
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	pruneCmd.Flags().StringVar(&lFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
}
