// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Config struct {
	DestDir         string `json:"destination"`
	PackageTemplate string `json:"pkgpagetmpl"`
	IndexTemplate   string `json:"pkgindextmpl"`
	DirIndex        string `json:"directory_index"`
}

type PkgCfg struct {
	Name       string `json:"name"`
	ImportPath string `json:"importpath"`
	VCS        string `json:"vcs"`
	Repo       string `json:"repo"`
	DocURL     string `json:"docurl"`
	Godoc      string `json:"godoc"`
	Descr      string `json:"description"`
	Humans     bool   `json:"human_readable"`
}
