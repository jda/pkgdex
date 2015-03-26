// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
)

// make package page
func genIndexPage(ps []PkgCfg) (err error) {
	outFile := filepath.Join(cfg.DestDir, cfg.DirIndex)
	out, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, defaultFileMode)
	if err != nil {
		return err
	}

	oc := IndexVars{cfg, ps}
	err = indexPage.Execute(out, oc)
	if err != nil {
		return err
	}

	return nil
}
