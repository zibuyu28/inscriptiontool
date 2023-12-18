/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"inscriptiontool/wallet"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore wallet",
	Long:  `restore wallet by private key, wallet info will be saved in a file named {name}.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return wallet.Restore(name, privateKey)
	},
}

var privateKey string

func init() {
	walletCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	restoreCmd.PersistentFlags().StringVarP(&privateKey, "private_key", "p", "", "wallet private key")
	restoreCmd.MarkPersistentFlagRequired("private_key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
