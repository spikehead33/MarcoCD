/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"marcocd/pkg/applications"
	"marcocd/pkg/infras/manifest_reader"

	"github.com/spf13/cobra"
)

type renderFlags struct {
	manfestPath string
}

var renFlags renderFlags

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render a module with given values",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		templateRenderer := applications.NewModuleTemplateRenderer(
			renFlags.manfestPath,
			manifest_reader.NewModuleManifestReader(),
		)

		jobSpecs, err := templateRenderer.Render()
		if err != nil {
			panic(err)
		}

		for _, jobSpec := range jobSpecs {
			fmt.Println(jobSpec)
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringVar(&renFlags.manfestPath, "manfestPath", "marcocd.yaml", "module root path")
}
