// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "html/template"

type IndexVars struct {
	Cfg  Config
	Pkgs []PkgCfg
}

// standard index page
var indexPageHTML string = `<html>
<head>
<meta name="generator" content="pkgdex">
<title>{{ if .Cfg.Title }}{{ .Cfg.Title }}{{ else }}go package index{{ end }}</title>
</head>
<body>
<h1>{{ if .Cfg.Title }}{{ .Cfg.Title }}{{ else }}go package index{{ end }}</h1>

<table border="1">
<tr>
<th>Name</th>
<th>Description</th>
<th>Import Path</th>
<th>Godoc</th>
</tr>
{{ range .Pkgs }}
 {{ if .NoIndex }}{{ else }}
 <tr>
 <td>{{ if .Humans }}<a href="{{ .Name }}">{{ .Name }}</a>{{ else }}{{ .Name }}{{ end }}</td>
 <td>{{ .Descr }}</td>
 <td>{{ .ImportPath }}</td>
 <td><a href="{{ .Godoc }}"><img src="{{ .Godoc }}?status.svg" alt="GoDoc"></a></td>
 </tr>
 {{ end }}
{{ end }}
</table>
<hr>
<i>Generated by <a href="http://github.com/jda/pkgdex">pkgdex</a></i>
</body>
</html>
`

var indexPage = template.Must(template.New("indexPage").Parse(indexPageHTML))
