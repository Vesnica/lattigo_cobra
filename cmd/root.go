// Copyright Vesnica
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// var key_path = "./key.toml"
// var data_path = "./data.toml"
// var stark_path = "/home/ubuntu/rust/stark-he/stark.toml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lattigo_cobra",
	Short: "A HE application base on lattigo",
	Long: `This program is a part of the STARK+HE prototype system.
Which is responsible for encrypting plaintext input using HE algorithm 
and decrypting the calculation result provided by STARK prover`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lattigo_cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
