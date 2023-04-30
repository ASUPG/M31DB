package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func dbengine(args []string) string {
	var returnval string = ""
	switch args[1] {
	case "select":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := strings.Replace(args[0], "\\main.exe", "\\db\\", -1)
		file = file + args[2] + ".json"

		type readResult struct {
			data []byte
			err  error
		}

		readCh := make(chan readResult)
		go func() {
			data, err := os.ReadFile(file)
			readCh <- readResult{data: data, err: err}
		}()

		select {
		case readRes := <-readCh:
			if readRes.err != nil {
				fmt.Println("Error Fetching Data:", readRes.err)
			} else {
				returnval = string(readRes.data)
			}
		}
	case "insert":
		args[2] = strings.ReplaceAll(args[2], "/", "\\")
		file := strings.Replace(args[0], "\\main.exe", "\\db\\", -1)
		file = file + args[2] + ".json"

		type writeResult struct {
			err error
		}

		writeCh := make(chan writeResult)

		poolSize := 5 // number of workers in the pool
		jobsCh := make(chan string, poolSize)
		for i := 0; i < poolSize; i++ {
			go func() {
				for job := range jobsCh {
					datab, err := os.ReadFile(file)
					ferr(err)
					mdata := formatJSON(string(datab))
					err = os.WriteFile(file, []byte(mdata), 0644)
					ferr(err)

					data, err := os.ReadFile(file)
					if err != nil {
						writeCh <- writeResult{err: err}
						return
					}
					argsmod := convandrotojson(job)
					mddata := string(data)
					modifieddata := strings.ReplaceAll(mddata, "]", "") + ",\n" + argsmod + "]"
					err = os.WriteFile(file, []byte(modifieddata), 0644)
					if err != nil {
						if strings.Contains(err.Error(), "file") {
							writeCh <- writeResult{err: errors.New("database does not exists")}
						} else {
							writeCh <- writeResult{err: err}
						}
						return
					}
				}
			}()
		}

		for i := 0; i < len(args)-3; i++ {
			jobsCh <- args[i+3]
		}
		close(jobsCh)

		var writeErr error
		for i := 0; i < len(args)-3; i++ {
			if res := <-writeCh; res.err != nil {
				writeErr = res.err
				break
			}
		}
		if writeErr != nil {
			fmt.Println("Error Inserting Data", writeErr)
		} else {
			returnval = "Done Successfully"
		}
	case "create":
		if args[2] != "table" {
			err := os.Mkdir("db\\"+args[2], 0755)
			if err == nil {
				returnval = "Done Successfully"
			} else {
				fmt.Println("Database Already exists")
			}
		}

	}
	return returnval
}
