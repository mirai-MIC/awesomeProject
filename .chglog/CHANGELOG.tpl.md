{{ if .Versions -}}
{{ range .Versions -}}

## {{ .Tag.Name }} ({{ datetime "2006-01-02" .Tag.Date }})

{{ if .CommitGroups -}}
{{ range .CommitGroups -}}

### {{ .Title }}

{{ range .Commits -}}

- {{ .Subject }} ({{ .Hash.Short }})
  {{ end }}
  {{ end }}
  {{ end }}
  {{ end -}}
  {{ end -}}
