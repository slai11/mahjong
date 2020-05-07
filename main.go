package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	m := map[string]*GameHolder{
		"test": &GameHolder{g: NewGame("test")},
	}

	server := &Server{
		Server: http.Server{
			Addr: ":80",
		},
	}

	if err := server.Start(m); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
