/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	nomad "github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list modules",
	Long:  `list the modules that is managed by MarcoCD`,
	Run: func(cmd *cobra.Command, args []string) {
		nc, err := nomad.NewClient(nomad.DefaultConfig())
		if err != nil {
			panic(err)
		}

		jobHandler := nc.Jobs()

		stubs, _, err := jobHandler.List(&nomad.QueryOptions{
			AllowStale: true,
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(formateModuleListOutput(stubs))
	},
}

func formateModuleListOutput(stubs []*nomad.JobListStub) string {
	lines := [][]string{}
	lines = append(lines, []string{"module", "jobID", "version"})

	for _, stub := range stubs {
		managedBy, ok := stub.Meta["managed-by"]
		if !ok || managedBy != "marcocd" {
			continue
		}

		jobID := stub.ID
		moduleName, ok := stub.Meta["module"]
		if !ok {
			moduleName = "None"
		}
		version, ok := stub.Meta["version"]
		if !ok {
			version = "None"
		}
		lines = append(lines, []string{moduleName, jobID, version})
	}

	// TODO: spaces alignment between each columne
	return ""
}

func init() {
	rootCmd.AddCommand(listCmd)
}
