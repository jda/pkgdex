// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// pretty print package summary
func printPkg(p PkgCfg) {
	fmt.Printf("%s (%s)\n", p.Name, p.ImportPath)
	fmt.Println("================================================")
	fmt.Println(p.Descr)
	fmt.Printf("Godoc: %s\n", p.Godoc)
	fmt.Println()
}
