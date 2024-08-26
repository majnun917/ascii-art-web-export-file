package main

import (
	"flag"
	"fmt"
	"net/http"

	f "ascii-web/handler"
)

func main() {
	// Serve static files with custom handler to prevent directory listing
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", f.NoDirListing(http.StripPrefix("/static/", fs)))

	// Handle main page requests
	http.HandleFunc("/", f.PathCheck("/", f.MainPage))

	// Handle form submission and result download
	http.HandleFunc("/submit", f.PathCheck("/submit", f.SubmitForm))
	http.HandleFunc("/download", f.PathCheck("/download", f.DownloadArt))

	// Parse the server port from the command line flags
	serverPort := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	fmt.Printf("http://localhost:%s - Server started on port\n", *serverPort)

	// Start the server && listen the requests
	http.ListenAndServe(":"+*serverPort, nil)
}
