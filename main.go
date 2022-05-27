package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

type AppParameters struct {
	version string
	bump    string
}

var parameters AppParameters

func main() {

	flag.StringVar(&parameters.version, "version", "", "specify the current version")
	flag.StringVar(&parameters.version, "v", "", "see -version")
	flag.StringVar(&parameters.bump, "bump", "patch", "specify which version to update")
	flag.StringVar(&parameters.bump, "b", "patch", "see -bump")
	flag.Parse()

	log.Debug(parameters)

	if parameters.version == "" {
		log.Fatal("no version specified")
		os.Exit(1)
	}
}
