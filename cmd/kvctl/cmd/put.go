// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/helinwang/kv"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put puts value into the kv service",
	Long: `put puts value into the kv service. For example:

put service-addr key value`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Printf("wrong number of arguments: %d, need: 3\n", len(args))
			return
		}

		c, err := kv.New(args[0])
		if err != nil {
			panic(err)
		}

		err = c.Put([]byte(args[1]), []byte(args[2]))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
