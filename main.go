package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Default directory for current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// We will server the files from the "files" directory
	dir = filepath.Join(dir, "files")
	fmt.Println(dir)

	// Get abspath for clearer logging
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}


	fs := http.FileServer(http.Dir(dir))

	http.Handle("/", fs)
	port := "8080"
	
	fmt.Printf("🚀 File server running on http://localhost:%s\n", port)
	fmt.Printf("📁 Serving files from: %s\n", absDir)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}