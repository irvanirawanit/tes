package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// run file sh
	cmd := exec.Command("sh", "file.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error file sh nya : ", err)
		log.Fatal(err)
	}

	// run server on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go infoServer()
		fmt.Fprintf(w, "Hello, World!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}

func infoServer() {

	fmt.Println("Machine/Server Information:")
	fmt.Println("===========================")

	// Retrieve and print the operating system details
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Compiler: %s\n", runtime.Compiler)
	fmt.Printf("Go Version: %s\n", runtime.Version())
	// check ubuntu or centos
	// check if ubuntu
	if runtime.GOOS == "linux" {
		// check if centos
		if _, err := os.Stat("/etc/centos-release"); err == nil {
			fmt.Printf("CentOS Version: %s\n", getCentOSVersion())
		} else {
			fmt.Printf("Ubuntu Version: %s\n", getUbuntuVersion())
		}
	}

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

func getCentOSVersion() string {
	// Read the version file
	version, err := os.ReadFile("/etc/centos-release")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(version))
}

func getUbuntuVersion() string {
	// Read the version file
	version, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	// Find the version line
	for _, line := range strings.Split(string(version), "\n") {
		if strings.HasPrefix(line, "VERSION=") {
			return strings.Trim(line, "VERSION=\"")
		}
	}
	return ""
}

func installLibreOffice() {
	cmd := exec.Command("sudo", "apt-get", "update")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("sudo", "apt-get", "install", "libreoffice", "-y")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("LibreOffice installation completed successfully.")
}
