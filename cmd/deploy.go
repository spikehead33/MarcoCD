/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"marcocd/pkg/applications"
	"marcocd/pkg/domains"

	nomad "github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

type deployFlags struct {
	manifestPath string
}

var dFlags deployFlags

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a marcocd module",
	Long:  `Deploy a marcocd into nomad cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		moduleManifest, err := domains.NewModuleManifestFromFile(
			dFlags.manifestPath,
		)
		if err != nil {
			panic(err)
		}

		renderer := applications.NewModuleTemplateRenderer(moduleManifest)

		nc, err := nomad.NewClient(nomad.DefaultConfig())
		if err != nil {
			panic(err)
		}

		deployer := applications.NewDeployer(moduleManifest.Name, nc, renderer)
		if err = deployer.Deploy(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	renderCmd.Flags().StringVar(&dFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
}
