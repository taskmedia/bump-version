package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

type AppParameters struct {
	version string
	bump    string
}

type Version struct {
	major int
	minor int
	patch int
}

var parameters AppParameters
var useVersionPrefix = false

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

	// if parameters.version starts with character v remove it
	if parameters.version[0] == 'v' {
		parameters.version = parameters.version[1:]
		useVersionPrefix = true
	}

	v := splitVersion(parameters.version)

	fmt.Println(v)
}

func splitVersion(version string) Version {
	versionParts := strings.Split(version, ".")

	if len(versionParts) != 3 {
		log.Fatal("invalid version format - expected 3 integers (e.g. v1.2.3)")
		os.Exit(1)
	}

	major, err := strconv.Atoi(versionParts[0])
	if err != nil {
		log.Fatal("could not parse major version - expected integer")
		log.Fatal(err)
		os.Exit(1)
	}
	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		log.Fatal("could not parse minor version - expected integer")
		log.Fatal(err)
		os.Exit(1)
	}
	patch, err := strconv.Atoi(versionParts[2])
	if err != nil {
		log.Fatal("could not parse patch version - expected integer")
		log.Fatal(err)
		os.Exit(1)
	}

	v := Version{major: major, minor: minor, patch: patch}

	return v
}
