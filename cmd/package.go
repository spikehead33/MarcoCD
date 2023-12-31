/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"marcocd/pkg/applications"
	"marcocd/pkg/domains"
	"marcocd/pkg/infras/tar_executor"

	"github.com/spf13/cobra"
)

type packageFlags struct {
	version      string
	output       string
	manifestPath string
}

var pFlags packageFlags

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package a macrocd module",
	Long:  `package a marcocd module into a zipped Tarball`,
	Run: func(cmd *cobra.Command, args []string) {
		moduleManifest, err := domains.NewModuleManifestFromFile(
			pFlags.manifestPath,
		)
		if err != nil {
			panic(err)
		}

		packager := applications.NewPackager(
			pFlags.manifestPath,
			pFlags.output,
			pFlags.version,
			moduleManifest,
			tar_executor.TarExecutor{},
		)
		if err := packager.Package(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVar(&pFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
	packageCmd.Flags().StringVar(&pFlags.version, "version", "0.0.0", "package version")
	packageCmd.Flags().StringVar(&pFlags.output, "name", "", "package name")
}
