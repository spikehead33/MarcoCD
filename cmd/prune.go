/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type pruneFlags struct {
	manifestPath string
}

var prFlags pruneFlags

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prune called")
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
	pruneCmd.Flags().StringVar(&prFlags.manifestPath, "manifestPath", "marcocd.yaml", "module root path")
}
