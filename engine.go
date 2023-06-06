package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func dbengine(args []string) string {
	var returnval string = ""
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	dbloc := dir + "\\db\\"
	switch args[1] {
	case "select":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := dbloc
		file = file + args[2] + ".json"
		dataCh := make(chan []byte)
		errCh := make(chan error)
		go func() {
			datab, err := os.ReadFile(file)
			if err != nil {
				errCh <- err
				return
			} else {
				dataCh <- datab
				err = os.WriteFile(file, datab, 0644)
				if err != nil {
					errCh <- err
				} else {
					errCh <- nil
				}
			}
		}()
		select {
		case data := <-dataCh:
			returnval = string(data)
		case err := <-errCh:
			fmt.Println("Error Fetching Data:", err)
		}
	case "insert":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := dbloc
		file = file + args[2] + ".json"

		dataCh := make(chan []byte)
		errCh := make(chan error)
		go func() {
			data, err := os.ReadFile(file)
			if err != nil {
				errCh <- err
				return
			}
			var argsmod string
			if args[3] != "json" {
				argsmod = convandrotojson(args[3])
				fmt.Println(argsmod)
			} else if args[3] == "json" {
				argsmod = formatJSON(args[4])
			}
			mddata := string(data)
			modifieddata := strings.ReplaceAll(mddata, "]", "") + ",\n" + argsmod + "]"
			dataCh <- []byte(formatJSON(modifieddata))
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
		if args[2] == "cluster" {
			err := os.Mkdir("db\\"+args[3], 0755)
			if err == nil {
				returnval = "Done Successfuly"
			} else {
				fmt.Println("Database Already exists")
			}
		} else if args[2] == "star" {
			filename := "db\\" + strings.Replace(args[3], "/", "\\", -1) + ".json"
			args[3] = strings.ReplaceAll(args[3], "\\", "/")
			filename = strings.ReplaceAll(filename, "\\", "/")
			data := formatJSON("[{\"name\":\"" + args[3] + "\"" + "," + "\"location\":" + "\"" + filename + "\"" + "}]")
			data = strings.ReplaceAll(data, "\\", "/")
			err := os.WriteFile(filename, []byte(data), 0644)
			if err != nil {
				returnval = ferr(err)
			} else {
				returnval = "Done Successfuly"
			}

		} else {
			fmt.Println(string("\033[31m"), "Error:Can't recognize what you are trying to create!", string("\033[0m"))
		}

	case "delete":
		file = strings.ReplaceAll(args[2], "/", "\\")
		file = dbloc + file + ".json"
		var csvtojson string
		var data []map[string]string
		jsontoarr := make(map[string]string)
		go func() {
			if args[3] == "json" {
				fmt.Println("json: ", true)
			} else if args[3] != "json" {
				csvtojson = formatJSON(convandrotojson(args[3]))
				databytes, err := os.ReadFile(file)
				returnval = ferr(err)
				go func() {
					done := make(chan bool)
					err = json.Unmarshal(databytes, &data)
					returnval = ferr(err)
					<-done
				}()
				err = json.Unmarshal([]byte(csvtojson), &jsontoarr)
				returnval = ferr(err)
				data = append(data, jsontoarr)
				fmt.Println(data)
			}
		}()
	}
	return returnval
}
