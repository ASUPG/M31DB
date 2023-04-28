package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fmt.Println(dbengine(args))
}

func dbengine(args []string) string {
	var returnval string = ""
	switch args[1] {
	case "select":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := strings.Replace(args[0], "\\main.exe", "\\db\\", -1)

		dataCh := make(chan []byte)
		errCh := make(chan error)

		go func() {
			data, err := os.ReadFile(file + args[2] + ".json")
			if err != nil {
				errCh <- err
				return
			}
			dataCh <- data
		}()

		select {
		case data := <-dataCh:
			returnval = string(data)
		case err := <-errCh:
			fmt.Println("Error Fetching Data:", err)
		}
	case "insert":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := strings.Replace(args[0], "\\main.exe", "\\db\\", -1)
		file = file + args[2] + ".json"

		dataCh := make(chan []byte)
		errCh := make(chan error)

		go func() {
			data, err := os.ReadFile(file)
			if err != nil {
				errCh <- err
				return
			}
			argsmod := convandrotojson(args[3])
			modifieddata := "[" + string(data) + ",\n" + argsmod + "]"
			dataCh <- []byte(modifieddata)
		}()

		go func() {
			err := os.WriteFile(file, <-dataCh, 00644)
			if err != nil {
				errCh <- err
				return
			}
			errCh <- nil
		}()

		select {
		case err := <-errCh:
			if err == nil {
				returnval = "Done Successfuly"
			} else {
				fmt.Println("Error Inserting Data", err)
			}
		}
	case "create":
		if args[2] != "table" {
			err := os.Mkdir("db\\"+args[2], 0755)
			if err == nil {
				returnval = "Done Successfuly"
			} else {
				fmt.Println("Database Already exists")
			}
		}

	}
	return returnval
}
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

// func formatJSON(input string) (string, error) {
// 	var out bytes.Buffer
// 	err := json.Indent(&out, []byte(input), "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return out.String(), nil
// }
