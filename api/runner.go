package api

import (
	"atlas/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/olekukonko/tablewriter"
)

// ExecuteActions displays menu and runs selected action
func ExecuteActions(client *auth.APIClient, af *ActionFile) {
	// Build menu
	options := make([]string, 0, len(af.Actions))
	for k, a := range af.Actions {
		options = append(options, fmt.Sprintf("%s - %s", k, a.Description))
	}
	options = append(options, "Exit")

	for {
		var choice string
		survey.AskOne(&survey.Select{
			Message: "Select an action:",
			Options: options,
		}, &choice)

		if choice == "Exit" {
			return
		}

		// Extract key from selection
		key := strings.Split(choice, " ")[0]
		action := af.Actions[key]
		runAction(client, action)
	}
}

// Run a single APIAction
func runAction(client *auth.APIClient, action APIAction) {
	// Fill in missing body values interactively
	if action.Body != nil {
		for k, v := range action.Body {
			if str, ok := v.(string); ok && str == "" {
				var input string
				survey.AskOne(&survey.Input{
					Message: fmt.Sprintf("Enter value for %s:", k),
				}, &input)
				action.Body[k] = input
			}
		}
	}

	var reqBody io.Reader
	if action.Body != nil {
		data, _ := json.Marshal(action.Body)
		reqBody = bytes.NewBuffer(data)
	}

	resp, err := client.DoRequest(action.Method, action.Path, &http.Request{
		Method: action.Method,
		Body:   io.NopCloser(reqBody),
	})
	if err != nil {
		fmt.Println("API request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var jsonResp interface{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		fmt.Println("Raw response:", string(body))
		return
	}

	// If array of objects → show table
	if arr, ok := jsonResp.([]interface{}); ok && len(arr) > 0 {
		printTable(arr)
	} else {
		pretty, _ := json.MarshalIndent(jsonResp, "", "  ")
		fmt.Println(string(pretty))
	}
}

// Print table from array of objects
func printTable(arr []interface{}) {
	table := tablewriter.NewWriter(os.Stdout)

	// Get headers from first item
	first := arr[0].(map[string]interface{})
	headers := make([]string, 0, len(first))
	for k := range first {
		headers = append(headers, k)
	}
	table.Header(headers)

	for _, item := range arr {
		row := make([]string, len(headers))
		m := item.(map[string]interface{})
		for i, h := range headers {
			row[i] = fmt.Sprintf("%v", m[h])
		}
		table.Append(row)
	}
	table.Render()
}
