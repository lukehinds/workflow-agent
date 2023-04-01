/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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

		if os.Getenv(RequestTokenEnvKey) == "" {
			fmt.Println("Error: Request token environment variable is not set")
			os.Exit(1)
		}

		requestURL := os.Getenv(RequestURLEnvKey)

		if requestURL == "" {
			fmt.Println("Error: Request URL environment variable is not set")
			os.Exit(1)
		}

		fmt.Println("Request URL is : " + requestURL)

		// Parse the access token
		tokenString := os.Getenv(RequestTokenEnvKey)
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			fmt.Println("Error: Unable to parse access token:", err)
			os.Exit(1)
		}

		// Extract the desired claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Error: Unable to extract claims from access token")
			os.Exit(1)
		}

		repository, ok := claims["repository"].(string)
		if !ok {
			fmt.Println("Error: Unable to extract repository claim")
			os.Exit(1)
		}

		jobWorkflowRef, ok := claims["job_workflow_ref"].(string)
		if !ok {
			fmt.Println("Error: Unable to extract job_workflow_ref claim")
			os.Exit(1)
		}

		fmt.Println("Repository: " + repository)
		fmt.Println("Job Workflow Ref: " + jobWorkflowRef)

		// Append audience to the URL
		url := requestURL + "&audience=sigstore"

		// Send the request with the access token
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error: Unable to create request:", err)
			os.Exit(1)
		}

		req.Header.Add("Authorization", "Bearer "+tokenString)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error: Unable to send request:", err)
			os.Exit(1)
		}

		var payload struct {
			Value string `json:"value"`
		}

		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&payload); err != nil {
			fmt.Println("Error: Unable to decode response body:", err)
			os.Exit(1)
		}

		fmt.Println("Payload Value: " + payload.Value)

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
