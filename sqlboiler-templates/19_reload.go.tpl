{{- $alias := .Aliases.Table .Table.Name -}}
{{- $schemaTable := .Table.Name | .SchemaTable -}}
{{- $canSoftDelete := .Table.CanSoftDelete }}

type {{$alias.UpSingular}}Reloader interface {
	Reload(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) error
	ReloadAll(o *{{$alias.UpSingular}}Slice, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) error
{{if .AddGlobal -}}
	ReloadG(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context{{end}}) error
	ReloadAllG(o *{{$alias.UpSingular}}Slice, {{if not .NoContext}}ctx context.Context{{end}}) error
{{end -}}
{{if .AddPanic -}}
	ReloadP(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}})
	ReloadAllP(o *{{$alias.UpSingular}}Slice, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}})
{{end -}}
{{if and .AddGlobal .AddPanic -}}
	ReloadGP(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context{{end}})
	ReloadAllGP(o *{{$alias.UpSingular}}Slice, {{if not .NoContext}}ctx context.Context{{end}})
{{end -}}
}

{{if .AddGlobal -}}
// ReloadG refetches the object from the database using the primary keys.
func (q {{$alias.UpSingular}}Query) ReloadG(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context{{end}}) error {
	if o == nil {
		return errors.New("{{.PkgName}}: no {{$alias.UpSingular}} provided for reload")
	}

	return q.Reload(o, {{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}})
}

{{end -}}

{{if .AddPanic -}}
// ReloadP refetches the object from the database with an executor. Panics on error.
func (q {{$alias.UpSingular}}Query) ReloadP(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) {
	if err := o.Reload({{if not .NoContext}}ctx, {{end -}} exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

{{if and .AddGlobal .AddPanic -}}
// ReloadGP refetches the object from the database and panics on error.
func (q {{$alias.UpSingular}}Query) ReloadGP(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context{{end}}) {
	if err := o.Reload({{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}}); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q {{$alias.UpSingular}}Query) Reload(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) error {
	ret, err := Find{{$alias.UpSingular}}({{if not .NoContext}}ctx, {{end -}} exec, {{.Table.PKey.Columns | stringMap (aliasCols $alias) | prefixStringSlice "o." | join ", "}})
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

{{if .AddGlobal -}}
// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q {{$alias.UpSingular}}Query) ReloadAllG(o *{{$alias.UpSingular}}Slice, {{if not .NoContext}}ctx context.Context{{end}}) error {
	if o == nil {
		return errors.New("{{.PkgName}}: empty {{$alias.UpSingular}}Slice provided for reload all")
	}

	return q.ReloadAll(o, {{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}})
}

{{end -}}

{{if .AddPanic -}}
// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (q {{$alias.UpSingular}}Query) ReloadAllP(o *{{$alias.UpSingular}}Slice, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) {
	if err := o.ReloadAll({{if not .NoContext}}ctx, {{end -}} exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

{{if and .AddGlobal .AddPanic -}}
// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (q {{$alias.UpSingular}}Query) ReloadAllGP(o *{{$alias.UpSingular}}Slice, {{if not .NoContext}}ctx context.Context{{end}}) {
	if err := o.ReloadAll({{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}}); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q {{$alias.UpSingular}}Query) ReloadAll(o *{{$alias.UpSingular}}Slice, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := {{$alias.UpSingular}}Slice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), {{$alias.DownSingular}}PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT {{$schemaTable}}.* FROM {{$schemaTable}} WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), {{if .Dialect.UseIndexPlaceholders}}1{{else}}0{{end}}, {{$alias.DownSingular}}PrimaryKeyColumns, len(*o)){{if and .AddSoftDeletes $canSoftDelete}} +
		"and {{"deleted_at" | $.Quotes}} is null"
		{{- end}}

	query := queries.Raw(sql, args...)

	err := query.Bind({{if .NoContext}}nil{{else}}ctx{{end}}, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "{{.PkgName}}: unable to reload all in {{$alias.UpSingular}}Slice")
	}

	*o = slice

	return nil
}
