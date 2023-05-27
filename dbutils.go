// Do Not delete this file or any code in it
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// For Fetching error for DRY code
func ferr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

// Functions for conversion of csv to json
func convandrotojson(data2 string) string {
	json := strings.Replace(data2, ",", ",\n", -1)
	json = "{\n" + json + "\n}"

	// Replace all occurrences of single quotes with double quotes
	data := json

	// Remove all spaces and newlines
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "\n", "")

	// Remove braces and split by commas
	pairs := strings.Split(data[1:len(data)-1], ",")
	m := make(map[string]string)

	// Create map from key-value pairs
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		m[kv[0]] = kv[1]
	}

	// Convert map to JSON string
	jsonString := "{"
	for k, v := range m {
		jsonString += "\"" + k + "\":\"" + v + "\","
	}
	jsonString = jsonString[:len(jsonString)-1] + "}"

	return jsonString
}

// Function for formating json string
func formatJSON(input string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	if err != nil {
		return ferr(err)
	}
	return out.String()
}

// Function for asking input works like nust the python function
func db_input(input string) string {
	fmt.Printf("%s", input)
	var data string
	fmt.Scanln(&data)
	return data
}

// Function for DB Initialization
func db_init(usr string, pwd string, wk string) {
	err := os.Mkdir("db", 0755)
	if err == nil {
		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println(string("\033[31m"), "Error While Creating config.json", err, string("\033[0m"))
		}
		config := make(map[string]string)
		config["username"] = usr
		config["password"] = pwd
		config["wk"] = wk
		configjson, err := json.Marshal(config)
		fmt.Println(ferr(err))
		file.Write([]byte(configjson))
		file.Close()
	} else {
		fmt.Println(string("\033[31"), err, string("\033[0m"))
	}
}

// function for removing a value from an slice
func db_rem(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
