// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// basic page is for redirect-to-godoc-only pages
var basicPackagePageHTML string = `<head>
<meta http-equiv="refresh" content="0; URL='{{ .Godoc }}'">
<meta name="go-import" content="{{ .ImportPath }} {{ .VCS }} {{ .Repo }}">
</head>
`
var basicPackagePage = template.Must(template.New("basicPackagePage").Parse(basicPackagePageHTML))

// detail page is human readable
var detailPackagePageHTML string = `

`
var detailPackagePage = template.Must(template.New("detailPackagePage").Parse(detailPackagePageHTML))

// make package page
func genPkgPage(p PkgCfg) (err error) {
	outDir := filepath.Join(cfg.DestDir, importName(p.ImportPath))
	err = os.Mkdir(outDir, defaultDirMode)
	if err != nil {
		exists := fmt.Sprintf("mkdir %s: file exists", outDir)
		if err.Error() != exists {
			return err
		}
	}

	outFile := filepath.Join(outDir, cfg.DirIndex)
	out, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, defaultFileMode)
	if err != nil {
		return err
	}

	if p.Humans {
		err = detailPackagePage.Execute(out, p)
	} else {
		err = basicPackagePage.Execute(out, p)
	}
	if err != nil {
		return err
	}

	return nil
}
