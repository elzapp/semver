/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"

	"github.com/elzapp/semver/internal/semver"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses a semver into it's parts and outputs it as a JSON object",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sv, _ := semver.ParseSemver(args[0])
		svs, _ := json.MarshalIndent(sv, "", "  ")
		fmt.Println(string(svs))
	},
}

var bumpMinorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bumps the minor version up by one 1.1.1 -> 1.2.0",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bumped, _ := semver.BumpMinor(args[0])
		fmt.Println(bumped)
	},
}

var bumpPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bumps the patch version up by one 1.1.1 -> 1.1.2",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		bumped, _ := semver.BumpPatch(args[0])
		fmt.Println(bumped)
	},
}

var bumpPrerelCmd = &cobra.Command{
	Use:   "prerel",
	Short: "Bumps the patch version 1.1.1 -> 1.1.2-0 -> 1.1.2-1",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		bumped, _ := semver.BumpPrerel(args[0])
		fmt.Println(bumped)
	},
}

var bumpMajorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bumps the major version 1.1.1 -> 2.0.0",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		bumped, _ := semver.BumpMajor(args[0])
		fmt.Println(bumped)
	},
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "bump a version",
	Long:  "bump a version using one of the subcommands [major, minor, patch, prerel]",
}

func init() {
	rootCmd.AddCommand(parseCmd)
	bumpCmd.AddCommand(bumpMinorCmd)
	bumpCmd.AddCommand(bumpPatchCmd)
	bumpCmd.AddCommand(bumpPrerelCmd)
	bumpCmd.AddCommand(bumpMajorCmd)
	rootCmd.AddCommand(bumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
