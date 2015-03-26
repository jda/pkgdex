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
var debug bool
var defaultCfgFile string = "pkgdex-prefs.json"
var defaultFileMode os.FileMode = 0644
var defaultDirMode os.FileMode = 0755

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] SOURCE\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func importName(ipath string) (name string) {
	parts := strings.Split(ipath, "/")
	name = parts[len(parts)-1]
	return name
}

func main() {
	destDir := ""
	testMode := false
	cfgFileName := ""
	flag.StringVar(&destDir, "dest", "", "output location (overrides pkgdex-prefs if set)")
	flag.BoolVar(&testMode, "test", false, "do not create output, show information on screen")
	flag.StringVar(&cfgFileName, "cfg", "", "override config file path (default is pkgdex-prefs.json in SOURCE")
	flag.BoolVar(&debug, "debug", false, "debug mode")

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

	if debug {
		fmt.Fprintf(os.Stderr, "Config as parsed:\n%+v\n", cfg)
		fmt.Fprintf(os.Stderr, "Confg from flags:\nDest: %s\nTest: %s\nConfig: %s\n", destDir, testMode, cfgFileName)
	}

	// flag overrides cfg
	if destDir != "" {
		cfg.DestDir = destDir
	}

	// if we aren't in test mode, make sure we have a dest dir
	if testMode == false {

		// make sure we have a dest dir
		if cfg.DestDir == "" {
			fmt.Fprintf(os.Stderr, "no destination set in config file or flag\n")
			os.Exit(1)
		}
	}

	// sort out other config defaults
	if cfg.DirIndex == "" {
		cfg.DirIndex = "index.html"
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

		if filepath.Ext(fileName) != ".json" {
			if debug {
				fmt.Fprintf(os.Stderr, "Skipping file %s because not json\n", fileName)
			}
			continue // skip because pkg files must be .json
		}

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
			pkg.Name = importName(pkg.ImportPath)
		}

		if pkg.Godoc == "" { // no godoc url specified so make one from import path
			pkg.Godoc = "http://godoc.org/" + pkg.ImportPath
		}

		pkgs = append(pkgs, pkg)

	}

	// pretty print and bail if in test mode
	if testMode {
		for i := range pkgs {
			printPkg(pkgs[i])
		}
		os.Exit(0)
	}

	// gen detail files
	for i := range pkgs {
		err = genPkgPage(pkgs[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generating page for package %s: %s\n", pkgs[i].Name, err)
			os.Exit(1)
		}
	}

	// gen index
	if cfg.NoIndex == false {
		err = genIndexPage(pkgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generating index page: %s\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
