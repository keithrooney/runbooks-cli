# {{ .Title }}

## Details

{{ .Details }}

## Impact

{{ .Impact }}

## Mitigation
{{ range $index, $element := .Mitigation.Steps }}
{{ $index }}. {{ $element.Name }}

   ```text
   {{ $element.Command }}
   ```
{{ end }}