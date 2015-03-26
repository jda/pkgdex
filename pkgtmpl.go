// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "html/template"

// basic page is for redirect-to-godoc-only pages
var basicPackagePageHTML string = `<head>
<meta http-equiv="refresh" content="0; URL='{{ .Godoc }}'">
<meta name="go-import" content="{{ .ImportPath }} {{ .VCS }} {{ .Repo }}">
</head>
`
var basicPackagePage = template.Must(template.New("basicPackagePage").Parse(basicPackagePageHTML))

// detail page is human readable
var detailPackagePageHTML string = `<html>
<head>
<meta http-equiv="refresh" content="0; URL='{{ .Godoc }}'">
<meta name="go-import" content="{{ .ImportPath }} {{ .VCS }} {{ .Repo }}">
<title>{{ .Name }}{{ if .Descr }} - {{ .Descr }}{{ end }}</title>
</head>
<body>
<h1>{{ .Name }}: {{ .ImportPath }}</h1>
<p><i>{{ .Descr }}</i><br>
{{ .VCS }}: {{ .Repo }}</p>
<h2>Reference</h2>
<a href="{{ .Godoc }}"><img src="{{ .Godoc }}?status.svg" alt="GoDoc"></a><br>
{{ if .DocURL }}See also: <a href="{{ .DocURL }}">{{ .DocURL }}</a>{{ end }}
</body>
</html>
`
var detailPackagePage = template.Must(template.New("detailPackagePage").Parse(detailPackagePageHTML))
