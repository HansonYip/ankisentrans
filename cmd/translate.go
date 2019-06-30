/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/HansonYip/ankisentrans/core"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate <src> <dest>",
	Short: "A tool for sentence translation for Anki.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires source and destination file paths")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("translation start")
		src, err := filepath.Abs(args[0])
		if err != nil {
			log.Fatal("source file path: ", err)
		}
		dest, err := filepath.Abs(args[1])
		if err != nil {
			log.Fatal("destination file path: ", err)
		}
		core.Process(src, dest)
		log.Println("translation done")
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
}
