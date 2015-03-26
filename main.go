// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var cfg Config
var defaultCfgFile string = "pkgdex-prefs.json"

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] SOURCE\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func godocify(ip string) (url string) {
	url = "http://godoc.org/" + ip
	return url
}

func main() {
	destDir := ""
	testMode := false
	cfgFileName := ""
	flag.StringVar(&destDir, "dest", "", "output location (overrides pkgdex-prefs if set)")
	flag.BoolVar(&testMode, "test", false, "do not create output, show information on screen")
	flag.StringVar(&cfgFileName, "cfg", "", "override config file path (default is pkgdex-prefs.json in SOURCE")

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	sourceDir := args[0]

	// read config if specified, try to read config file in sourcedir
	// and die if config provided but invalid
	forceCfgFile := true
	if cfgFileName == "" {
		forceCfgFile = false
		cfgFileName = filepath.Join(sourceDir, defaultCfgFile)
	}

	configFile, err := ioutil.ReadFile(cfgFileName)
	if err != nil {
		if forceCfgFile == true {
			fmt.Fprintf(os.Stderr, "Could not open config file %s because %s\n", cfgFileName, err)
			os.Exit(1)
		}
	} else {
		err = json.Unmarshal(configFile, &cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse (json) config file %s because %s\n", cfgFileName, err)
			os.Exit(1)
		}
	}

	// now override config file with cli params
	if destDir != "" {
		cfg.DestDir = destDir
	}

	// walk sourcedir. not bothering with subdirs because one level is good enough for now?
	fileList, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open source directory %s because %s\n", sourceDir, err)
		os.Exit(1)
	}

	// process package defs
	pkgs := make([]PkgCfg, 0)
	for fileN := range fileList {
		if fileList[fileN].IsDir() {
			continue // skip because dir
		}
		if fileList[fileN].Name() == defaultCfgFile {
			continue // skip because deafult config file
		}

		fileName := filepath.Join(sourceDir, fileList[fileN].Name())
		pkgCfgFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open package config %s because %s\n", fileName, err)
			continue
		}

		pkg := PkgCfg{}
		err = json.Unmarshal(pkgCfgFile, &pkg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse package config (json) %s because %s\n", fileName, err)
			continue
		}

		// clean up pkg info
		if pkg.Name == "" { // no package name so set to last part of importpath
			importPathParts := strings.Split(pkg.ImportPath, "/")
			pkg.Name = importPathParts[len(importPathParts)-1]
		}

		pkgs = append(pkgs, pkg)

	}

	if testMode {
		for i := range pkgs {
			printPkg(pkgs[i])
		}
		os.Exit(0)
	}

	// try to open dir
	// check for global config file (pkgdex-prefs.json)
	// walk other files
	// gen index
	// gen detail files
}
