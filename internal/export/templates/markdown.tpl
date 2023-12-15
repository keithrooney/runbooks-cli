# {{ .Title }}

## Details

{{ .Details }}

## Impact

{{ .Impact }}

## Mitigation

### When

To execute the mitigation steps, the condition below must evaluate to `true` first.

```
{{ .Mitigation.Assertion }}
```

### Steps

Below are the steps that must be performed in order to mitigate this incident.
{{ range $index, $element := .Mitigation.Steps }}
{{ $index }}. {{ $element.Name }}

   ```text
   {{ $element.Shell.Command }}
   ```
{{ if $element.Shell.Assertion }}
   Verify the command completed as expected, asserting the below command returns `true` :-

    ```text
   {{- $string := printf "%s" $element.Shell.Assertion -}}
   {{- $lines := split $string "\n" -}}
   {{- range $line := $lines }}
   {{- printf "\n    %s" $line -}}
   {{- end -}}
   ```
{{ end }}
{{- end -}}