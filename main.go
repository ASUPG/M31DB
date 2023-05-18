package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type WorkerPool struct {
	// Number of workers in the pool.
	workerCount int

	// Mutex to protect the workerCount.
	workerCountMutex sync.Mutex

	// Channel to send work items to the workers.
	workChannel chan func()

	// WaitGroup to wait for all workers to finish.
	waitGroup sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with the given number of workers.
func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		workerCount: workerCount,
		workChannel: make(chan func()),
	}

	for i := 0; i < workerCount; i++ {
		go pool.worker()
	}

	return pool
}

// Submit submits a work item to the pool.
func (pool *WorkerPool) Submit(f func()) {
	pool.workChannel <- f
}

// Wait waits for all work items in the pool to finish.
func (pool *WorkerPool) Wait() {
	pool.waitGroup.Wait()
}

func (pool *WorkerPool) worker() {
	for f := range pool.workChannel {
		f()
	}

	pool.waitGroup.Done()
}
func main() {
	args := os.Args
	if args[1] == "start" {
		filecont, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Println(string("\033[31m"), "Please Initialize your config using 'm31 init", string("\033[0m"))
			os.Exit(1)
		}
		config := map[string]string{}
		err = json.Unmarshal(filecont, &config)
		if err != nil {
			fmt.Println(string("\033[31m"), "Please Initialize your config using 'm31 init", string("\033[0m"))
			os.Exit(1)
		}
		pooln, _ := strconv.ParseInt(config["wk"], 10, 64)
		pool := NewWorkerPool(int(pooln))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			options := r.URL.Query().Get("options")
			if options != "" {
				done := make(chan bool)
				pool.Submit(func() {
					optionList := strings.Split(options, ",")
					optionList = append([]string{""}, optionList...)
					result := dbengine(optionList)
					fmt.Fprintf(w, result)
					done <- true // send signal to wait for the worker to complete
				})
				<-done // wait for the worker to complete before returning the response
				return
			}
		})

		fmt.Println(string("\033[32m"), "Database listening on port 6787", string("\033[0m"))
		fmt.Println(string("\033[33m"), "Press CTRL + C to Stop", string("\033[0m"))
		http.ListenAndServe(":6787", nil)
	} else if args[1] == "init" {
		err := os.Mkdir("db", 0755)
		if err == nil {
			file, err := os.Create("config.json")
			if err != nil {
				fmt.Println(string("\033[31m"), "Error While Creating config.json", err, string("\033[0m"))
			}
			usr := input("Define a username: ")
			pwd := input("Define a password: ")
			wk := input("Worker Poool (Define in Numbers)\n(Hint: Higher it is the more concurrent connection it can handel but also use higher resources):")
			config := make(map[string]string)
			config["username"] = usr
			config["password"] = pwd
			config["wk"] = wk
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
