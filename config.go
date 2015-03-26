// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Config struct {
	DestDir         string `json:"destination"`
	PackageTemplate string `json:"pkgpagetmpl"`
	IndexTemplate   string `json:"pkgindextmpl"`
	DirIndex        string `json:"directory_index"`
	Title           string `json:"title"`
	OwnerContact    string `json:"owner_contact"`
	OwnerURL        string `json:"owner_url"`
	NoIndex         bool   `json:"noindex"`
}

type PkgCfg struct {
	Name       string `json:"name"`
	ImportPath string `json:"importpath"`
	VCS        string `json:"vcs"`
	Repo       string `json:"repo"`
	DocURL     string `json:"docurl"`
	Godoc      string `json:"godoc"`
	Descr      string `json:"description"`
	Humans     bool   `json:"humans"`
	NoIndex    bool   `json:"noindex"`
}
