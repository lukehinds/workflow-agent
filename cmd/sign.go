/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

const (
	RequestTokenEnvKey = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
	RequestURLEnvKey   = "ACTIONS_ID_TOKEN_REQUEST_URL"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if os.Getenv(RequestTokenEnvKey) != "" {
			fmt.Println("Request is : " + os.Getenv(RequestTokenEnvKey))
		}

		if os.Getenv(RequestURLEnvKey) != "" {
			fmt.Println("Request URL is : " + os.Getenv(RequestURLEnvKey))
			os.Exit(1)
		}

		url := os.Getenv(RequestURLEnvKey) + "&audience=aud"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Authorization", "Bearer "+os.Getenv(RequestTokenEnvKey))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		var payload struct {
			Value string `json:"value"`
		}

		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&payload); err != nil {
			fmt.Println(err)
		}

		fmt.Println(payload.Value)

	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
