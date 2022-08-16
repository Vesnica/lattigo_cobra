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

type Data struct {
	Modulus []uint64
	Values  [][][][]uint64
}

var inputA uint64
var inputB uint64
var inputC uint64
var data_path string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Generate some encrypted content",
	Long: `Use key.toml and lattigo HE algorithm to encrypt
some plain value, then output encrypted value to data.toml`,
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

		encoder := bfv.NewEncoder(params)
		sk := bfv.NewSecretKey(params)
		sk.UnmarshalBinary(key.Private)
		encryptor := bfv.NewEncryptor(params, sk)

		A := []uint64{inputA}
		B := []uint64{inputB}
		C := []uint64{inputC}
		fmt.Printf("A: %v\n", A[0])
		fmt.Printf("B: %v\n", B[0])
		fmt.Printf("C: %v\n", C[0])
		a_plus_b := (A[0] + B[0]) % paramDef.T
		var ret uint64
		if a_plus_b > C[0] {
			ret = a_plus_b - C[0]
		} else {
			ret = a_plus_b + paramDef.T - C[0]
		}
		fmt.Printf("A+B-C: %v\n", ret)

		APlain := bfv.NewPlaintext(params)
		encoder.Encode(A, APlain)
		ACipher := encryptor.EncryptNew(APlain)

		BPlain := bfv.NewPlaintext(params)
		encoder.Encode(B, BPlain)
		BCipher := encryptor.EncryptNew(BPlain)

		CPlain := bfv.NewPlaintext(params)
		encoder.Encode(C, CPlain)
		CCipher := encryptor.EncryptNew(CPlain)

		f, err := os.Create(data_path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()

		data := Data{
			params.RingQ().Modulus,
			[][][][]uint64{
				{ACipher.Value[0].Coeffs, ACipher.Value[1].Coeffs},
				{BCipher.Value[0].Coeffs, BCipher.Value[1].Coeffs},
				{CCipher.Value[0].Coeffs, CCipher.Value[1].Coeffs},
			},
		}
		if err := toml.NewEncoder(f).Encode(&data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encryptCmd.Flags().Uint64VarP(&inputA, "A", "a", 10000, "Input: A")
	encryptCmd.Flags().Uint64VarP(&inputB, "B", "b", 1000, "Input: B")
	encryptCmd.Flags().Uint64VarP(&inputC, "C", "c", 100, "Input: C")
	encryptCmd.Flags().StringVarP(&data_path, "data_file", "d", "./data.toml", "encrypted data file path")
}
