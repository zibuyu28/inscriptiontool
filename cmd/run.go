/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"inscriptiontool/run"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run inscription",
	Long:  `run inscription`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run.Run(i)
	},
}

var i = run.Info{
	Account:        "",
	Rpcs:           []string{},
	To:             "",
	Amount:         0,
	GF:             1,
	GasLimitFactor: 1,
	Data:           "",
	Count:          1,
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	runCmd.PersistentFlags().StringSliceVarP(&i.Rpcs, "rpcUrls", "r", []string{}, "rpc 地址,多地址用 *逗号* 隔开")
	runCmd.MarkPersistentFlagRequired("rpcUrl")
	runCmd.PersistentFlags().StringVarP(&i.Account, "account", "a", "", "账户名称")
	runCmd.MarkPersistentFlagRequired("account")
	runCmd.PersistentFlags().StringVarP(&i.To, "to", "t", "", "转账To地址, (默认是转给自己)")
	runCmd.PersistentFlags().Float64VarP(&i.Amount, "amount", "m", 0, "转账金额, (默认是0)")
	runCmd.PersistentFlags().Float64VarP(&i.GF, "gf", "g", 1, "gas倍数")
	runCmd.PersistentFlags().IntVarP(&i.GasLimitFactor, "glf", "l", 2, "gasLimit倍数, 默认是2，即21000*2")
	runCmd.PersistentFlags().StringVarP(&i.Data, "data", "d", "", "data数据,可以是json原文，也可以十六进制, data:,{xxxx}, 0x6137sdxxxx")
	runCmd.MarkPersistentFlagRequired("data")
	runCmd.PersistentFlags().IntVarP(&i.Count, "count", "c", 1, "重复次数")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
