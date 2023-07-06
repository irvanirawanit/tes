package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	// Print machine/server information
	infoServer()
	// run server on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go infoServer()
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
}

func infoServer() {

	fmt.Println("Machine/Server Information:")
	fmt.Println("===========================")

	// Retrieve and print the operating system details
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)

	// Retrieve and print the host name
	host, err := os.Hostname()
	if err == nil {
		fmt.Printf("Host Name: %s\n", host)
	}

	// Retrieve and print the number of CPUs
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)

	// Retrieve and print the Goroutine count
	numGoroutine := runtime.NumGoroutine()
	fmt.Printf("Number of Goroutines: %d\n", numGoroutine)

	// Retrieve and print memory statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Allocated Memory: %.2f MB\n", float64(memStats.Alloc)/(1024*1024))
	fmt.Printf("Total Allocated Memory: %.2f MB\n", float64(memStats.TotalAlloc)/(1024*1024))
	fmt.Printf("Heap Memory: %.2f MB\n", float64(memStats.HeapAlloc)/(1024*1024))
	fmt.Printf("Heap Objects: %d\n", memStats.HeapObjects)

	// Retrieve and print environment variables
	envVars := os.Environ()
	fmt.Println("\nEnvironment Variables:")
	fmt.Println("======================")
	for _, envVar := range envVars {
		envVarParts := strings.SplitN(envVar, "=", 2)
		fmt.Printf("%s = %s\n", envVarParts[0], envVarParts[1])
	}
}
