// Copyright 2015 Jonathan Auer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "html/template"

// basic page is for redirect-to-godoc-only pages
var packagePageHTML string = `{{if .Humans}}<html>
<head>
<meta name="go-import" content="{{ .ImportPath }} {{ .VCS }} {{ .Repo }}">
<title>{{ .Name }}{{ if .Descr }} - {{ .Descr }}{{ end }}</title>
</head>
<body>
<h1>{{ .Name }}</h1>
<p><i>{{ .Descr }}</i><br>
{{ .VCS }}: <a href="{{ .Repo }}">{{ .Repo }}</a></p>
<h2>Reference</h2>
<a href="{{ .Godoc }}"><img src="{{ .Godoc }}?status.svg" alt="GoDoc"></a><br>
{{ if .DocURL }}See also: <a href="{{ .DocURL }}">{{ .DocURL }}</a>{{ end }}
</body>
</html>{{ else }}<head>
<meta http-equiv="refresh" content="0; URL='{{ .Godoc }}'">
<meta name="go-import" content="{{ .ImportPath }} {{ .VCS }} {{ .Repo }}">
</head>{{ end }}
`
var packagePage = template.Must(template.New("packagePage").Parse(packagePageHTML))
