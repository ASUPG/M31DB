package main

import (
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
			returnval = ferr(err)
			mdata := formatJSON(string(datab))
			err = os.WriteFile(file, []byte(mdata), 0644)
			returnval = ferr(err)
		}()
		go func() {
			data, err := os.ReadFile(file)
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
		file := dbloc
		file = file + args[2] + ".json"

		dataCh := make(chan []byte)
		go func() {
			datab, err := os.ReadFile(file)
			returnval = ferr(err)
			mdata := formatJSON(string(datab))
			err = os.WriteFile(file, []byte(mdata), 0644)
			returnval = ferr(err)
		}()
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
			} else {
			}
			mddata := string(data)
			modifieddata := strings.ReplaceAll(mddata, "]", "") + ",\n" + argsmod + "]"
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
		if args[2] == "cluster" {
			err := os.Mkdir("db\\"+args[3], 0755)
			if err == nil {
				returnval = "Done Successfuly"
			} else {
				fmt.Println("Database Already exists")
			}
		} else if args[2] == "star" {
			filename := strings.Replace(args[3], "/", "\\", -1)
			file, err := os.Create("db\\" + filename + ".json")
			fmt.Println(ferr(err))
			n, err := file.Write([]byte(formatJSON("[{\"name\":\"" + args[3] + "\"}]")))
			print(n)
			fmt.Println(ferr(err))
			file.Close()
			fmt.Println("Done Successfuly")
		} else {
			fmt.Println(string("\033[31m"), "Error:Can't recognize what you are trying to create!", string("\033[0m"))
		}

		// case "delete":
		// 	file := strings.Replace(args[0], "\\main.exe", "\\db\\", -1)
		// 	file = file + args[2] + ".json"
		// 	filecontsup, err := os.ReadFile(file)
		// 	returnval = ferr(err)
		// 	nfilecontsup := string(filecontsup)
		// 	nfilecontsup = strings.Replace(nfilecontsup, "\n", "", -1)
		// 	fmt.Println(nfilecontsup)
	case "delete":
		file = strings.ReplaceAll(args[2], "/", "\\")
		file = dbloc + file + ".json"
		// Formatter
		go func() {
			datab, err := os.ReadFile(file)
			returnval = ferr(err)
			mdata := formatJSON(string(datab))
			err = os.WriteFile(file, []byte(mdata), 0644)
			returnval = ferr(err)
		}()
		go func() {

		}()
	}
	return returnval
}
