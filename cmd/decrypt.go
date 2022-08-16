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

type Stark struct {
	Result [][][]uint64 `toml:"result"`
	Proof  string       `toml:"proof"`
}

var proof_path string

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt result to plain value",
	Long: `Use key.toml and lattigo HE algorithm to decrypt the result
provided by STARK prover, which stored in stark.toml`,
	Run: func(_ *cobra.Command, _ []string) {
		paramDef := bfv.PN12QP101pq
		paramDef.T = paramDef.Q[0] // 0x800004001 = 34359754753

		params, err := bfv.NewParametersFromLiteral(paramDef)
		if err != nil {
			panic(err)
		}

		key := Key{}
		if _, err := toml.DecodeFile(key_path, &key); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		stark := Stark{}
		if _, err := toml.DecodeFile(proof_path, &stark); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		encoder := bfv.NewEncoder(params)
		sk := bfv.NewSecretKey(params)
		sk.UnmarshalBinary(key.Private)
		decryptor := bfv.NewDecryptor(params, sk)
		cipher := bfv.NewCiphertext(params, 1)

		cipher.Value[0].Coeffs = stark.Result[0]
		cipher.Value[1].Coeffs = stark.Result[1]

		result := encoder.DecodeUintNew(decryptor.DecryptNew(cipher))
		fmt.Printf("result: %v\n", result[0])
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decryptCmd.Flags().StringVarP(&proof_path, "proof_file", "p", "/home/ubuntu/rust/stark-he/stark.toml", "proof file path")
}
