package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	args := os.Args
	if args[1] != "start" {
		fmt.Println(dbengine(args))
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			options := r.URL.Query().Get("options")
			if options != "" {
				optionList := strings.Split(options, ",")
				result := dbengine(optionList)
				fmt.Println(result)
				fmt.Fprintf(w, "Result: %s", result)
			}
			fmt.Fprintf(w, "Hello, world!")
		})

		fmt.Println("Server listening on port 8080...")
		http.ListenAndServe(":8080", nil)
	}
	runtime.GC()
}
