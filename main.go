package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	args := os.Args
	if args[1] == "start" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			options := r.URL.Query().Get("options")
			// username := r.URL.Query().Get("username")
			// password := r.URL.Query().Get("password")
			if options != "" {
				optionList := strings.Split(options, ",")
				result := dbengine(optionList)
				fmt.Println(result)
				fmt.Fprintf(w, "Result: %s", result)
			}
			fmt.Fprintf(w, "Hello, world!")
		})

		fmt.Println("Server listening on port 6787...")
		http.ListenAndServe(":6787", nil)
	} else if args[1] == "init" {
		err := os.Mkdir("db", 0755)
		if err == nil {
			file, err := os.Create("config.json")
			if err != nil {
				fmt.Println("Error: ", err)
			}
			usr := input("Define a username: ")
			pwd := input("Define a password: ")
			config := make(map[string]string)
			config["username"] = usr
			config["password"] = pwd
			configjson, err := json.Marshal(config)
			fmt.Println(ferr(err))
			file.Write([]byte(configjson))
			file.Close()
		}
	} else {
		fmt.Println(dbengine(args))
	}
	runtime.GC()
}
