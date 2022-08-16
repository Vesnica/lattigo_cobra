// Copyright Vesnica
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package cmd

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/tuneinsight/lattigo/v3/bfv"
)

type Key struct {
	Private []byte
}

var key_path string

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Generate key pair",
	Long:  `Generate a new key pair, and save it to file`,
	Run: func(_ *cobra.Command, _ []string) {
		paramDef := bfv.PN12QP101pq
		paramDef.T = paramDef.Q[0] // 0x800004001 = 34359754753

		params, err := bfv.NewParametersFromLiteral(paramDef)
		if err != nil {
			panic(err)
		}

		kgen := bfv.NewKeyGenerator(params)
		sk, _ := kgen.GenKeyPair()
		skb, err := sk.MarshalBinary()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		data := Key{skb}

		f, err := os.Create(key_path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()

		if err := toml.NewEncoder(f).Encode(&data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	keyCmd.Flags().StringVarP(&key_path, "key_file", "k", "./key.toml", "key file path")
}
