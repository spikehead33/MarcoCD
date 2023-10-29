/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"marcocd/pkg"
	"marcocd/pkg/appcontroller"
	"marcocd/pkg/repomonitor"
	"marcocd/pkg/settings"

	nomad "github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

var serverSetting settings.ServerSettings

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "marcocd server",
	Long:  `MarcoCD server, a gitops server for nomad`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(serverSetting.GitHubAppID)
		// fmt.Println(serverSetting.GitHubAppInstallationID)
		// fmt.Println(serverSetting.GitHubAppPrivateKeyPath)

		nc, err := nomad.NewClient(nomad.DefaultConfig())
		if err != nil {
			panic(err)
		}

		appController := appcontroller.New(nc)
		repomonitor := repomonitor.New(
			serverSetting.GitHubAppID,
			serverSetting.GitHubAppInstallationID,
			serverSetting.GitHubAppPrivateKeyPath,
		)

		s := pkg.NewServer(appController, repomonitor)
		pkg.StartServer(":3000", s)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.Flags().Int64VarP(&serverSetting.GitHubAppID, "githubAppID", "g", 123, "GitHub App ID for github authentication")
	rootCmd.Flags().Int64VarP(&serverSetting.GitHubAppInstallationID, "githubAppInstallationID", "i", 456, "GitHub App Installation ID for github authentication")
	rootCmd.Flags().StringVarP(&serverSetting.GitHubAppPrivateKeyPath, "githubAppPrivateKeyPath", "k", "githubapp.pem", "GitHub App Private Key for github authentication")
	rootCmd.Flags().StringVarP(&serverSetting.GitOpsRepoPath, "repo", "p", "", "GitOps repository codebase path. remote path or local path are both acceptable")
}
