package main

import (
	"fmt"
	"net/http"

	"github.com/ccutch/goparcel"
)

func main() {
	worker, err := goparcel.Open("./connors-spot", "frontend/index.html")

	if err != nil && goparcel.IsNotFoundError(err) {
		fmt.Println("Setting up workspace")
		if err = worker.SetupWorkspace(); err != nil {
			panic(err)
		}
	}

	fmt.Println("Worker at: " + worker.Workspace())
	defer worker.Close()
	go worker.Start()

	http.Handle("/", worker.FileServer())
	http.ListenAndServe(":4000", nil)
}
