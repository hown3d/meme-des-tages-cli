package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("serving on :9090")
	err = http.ListenAndServe(":9090", http.FileServer(http.Dir(filepath.Join(cwd, "assets"))))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
