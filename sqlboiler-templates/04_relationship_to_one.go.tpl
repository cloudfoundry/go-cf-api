{{- if .Table.IsJoinTable -}}
{{- else -}}
	{{- range $fkey := .Table.FKeys -}}
		{{- $ltable := $.Aliases.Table $fkey.Table -}}
		{{- $ftable := $.Aliases.Table $fkey.ForeignTable -}}
		{{- $rel := $ltable.Relationship $fkey.Name -}}
		{{- $canSoftDelete := (getTable $.Tables $fkey.ForeignTable).CanSoftDelete }}
// {{$rel.Foreign}} pointed to by the foreign key.
func (q {{$ltable.DownSingular}}Query) {{$rel.Foreign}}(o *{{$ltable.UpSingular}}, mods ...qm.QueryMod) ({{$ftable.DownSingular}}Query) {
	queryMods := []qm.QueryMod{
		qm.Where("{{$fkey.ForeignColumn | $.Quotes}} = ?", o.{{$ltable.Column $fkey.Column}}),
		{{if and $.AddSoftDeletes $canSoftDelete -}}
		qmhelper.WhereIsNull("deleted_at"),
		{{- end}}
	}

	queryMods = append(queryMods, mods...)

	query := {{$ftable.UpPlural}}(queryMods...)
	queries.SetFrom(query.Query, "{{.ForeignTable | $.SchemaTable}}")

	return query
}
{{- end -}}
{{- end -}}
