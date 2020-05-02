package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	m := map[string]*Game{
		"test": NewGame("test"),
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
