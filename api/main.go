package main

import (
	"log"
	"net/http"

	"github.com/adrianosela/tagatree/api/config"
	"github.com/adrianosela/tagatree/api/ctrl"
)

var (
	// injected at build-time
	version string
)

func main() {
	ctrl, err := ctrl.NewController(config.FromEnv(version))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":80", ctrl.Router()); err != nil {
		log.Fatal(err)
	}
}
