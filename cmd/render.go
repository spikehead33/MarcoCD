/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"marcocd/pkg/applications"
	"marcocd/pkg/domains"

	"github.com/spf13/cobra"
)

type renderFlags struct {
	manifestPath string
}

var renFlags renderFlags

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render a module with given values",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		moduleManifest, err := domains.NewModuleManifestFromFile(
			renFlags.manifestPath,
		)
		if err != nil {
			panic(err)
		}

		renderer := applications.NewModuleTemplateRenderer(
			moduleManifest,
		)

		jobSpecs, err := renderer.Render()
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
	renderCmd.Flags().StringVar(&renFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
}
