/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"marcocd/pkg/applications"
	"marcocd/pkg/infras/manifest_reader"

	nomad "github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

type deployFlags struct {
	moduleName   string
	manifestPath string
}

var dFlags deployFlags

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a marcocd module",
	Long:  `Deploy a marcocd into nomad cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		renderer := applications.NewModuleTemplateRenderer(
			dFlags.manifestPath, manifest_reader.NewModuleManifestReader())

		nc, err := nomad.NewClient(nomad.DefaultConfig())
		if err != nil {
			panic(err)
		}

		deployer := applications.NewDeployer(nc, renderer)
		if err = deployer.Deploy(dFlags.moduleName); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVar(&dFlags.moduleName, "module", "", "module to be deploy to the nomad")
	renderCmd.Flags().StringVar(&dFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
}
