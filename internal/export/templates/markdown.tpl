# {{ .Title }}

## Details

{{ .Details }}

## Mitigation
{{ range $index, $element := .Mitigation.Steps }}
{{ $index }}. {{ $element.Name }}

   ```text
   {{ $element.Command }}
   ```
{{ end }}