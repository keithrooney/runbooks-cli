# {{ .Runbook.Title }}

## Details

{{ .Runbook.Details }}

## Impact

{{ .Runbook.Impact }}

## Mitigation

### When

To execute the mitigation steps, the condition below must evaluate to `true` first.

```text
{{ .Runbook.Mitigation.Clause }}
```

### Steps

Below are the steps that must be performed in order to mitigate this incident.
{{- printf "\n" -}}
{{- $formatOptions := .FormatOptions -}}
{{ range $index, $element := .Runbook.Mitigation.Steps }}
{{ printf "%d. %s\n\n" $index $element.Name -}}
{{ $element.Command.Format $formatOptions }}
{{- printf "\n" -}}
{{- end -}}
