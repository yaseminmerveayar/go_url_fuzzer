/*
Copyright © 2022 Yasemin Merve Ayar <yaseminmerve.ayar@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yaseminmerveayar/fuzzer/config"
	"github.com/yaseminmerveayar/fuzzer/requests"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fuzzer",
		Short: "Fuzz Url",
		Long:  `Fuzzer is a URL fuzzing tool which can help you find possible url paths with a wordlist`,

		Run: func(cmd *cobra.Command, args []string) {
			requests.Execute()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.AppFlag.Wordlist, "wordlist", "w", "", "path of the wordlist")
	rootCmd.MarkFlagRequired("wordlist")
	rootCmd.PersistentFlags().StringVarP(&config.AppFlag.URL, "url", "u", "", "path of the website url -- URL/FUZZ")
	rootCmd.MarkFlagRequired("url")
	rootCmd.PersistentFlags().StringVarP(&config.AppFlag.RequestType, "method", "m", config.DefaultRequestType, "request type -- GET, POST ")
	rootCmd.PersistentFlags().IntVarP(&config.AppFlag.StatusShow, "show", "s", config.DefaultStatusShow, "only show the status code you chooose -- for 404 change the -f flag")
	rootCmd.PersistentFlags().IntVarP(&config.AppFlag.StatusHide, "filter", "f", config.DefaultStatusHide, "hide the status code you chooose")
}
