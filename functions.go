package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

func ferr(err error) {
	if err != nil {
		print(err)
	}
}
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

func formatJSON(input string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	if err != nil {
		ferr(err)
	}
	return out.String()
}
