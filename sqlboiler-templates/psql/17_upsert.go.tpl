{{- $alias := .Aliases.Table .Table.Name}}
{{- $schemaTable := .Table.Name | .SchemaTable}}

type {{$alias.UpSingular}}Upserter interface {
	Upsert(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, updateColumns, insertColumns boil.Columns) error
{{if .AddGlobal -}}
	UpsertG(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context, {{end -}} updateColumns, insertColumns boil.Columns) error
{{end -}}
{{if and .AddGlobal .AddPanic -}}
	UpsertGP(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context, {{end -}} updateColumns, insertColumns boil.Columns)
{{end -}}
{{if .AddPanic -}}
	UpsertP(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, updateColumns, insertColumns boil.Columns)
{{end -}}
}

{{if .AddGlobal -}}
// UpsertG attempts an insert, and does an update or ignore on conflict.
func (q {{$alias.DownSingular}}Query) UpsertG(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context, {{end -}} updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert({{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}}, updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

{{end -}}

{{if and .AddGlobal .AddPanic -}}
// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (q {{$alias.DownSingular}}Query) UpsertGP(o *{{$alias.UpSingular}}, {{if not .NoContext}}ctx context.Context, {{end -}} updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert({{if .NoContext}}boil.GetDB(){{else}}ctx, boil.GetContextDB(){{end}}, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

{{if .AddPanic -}}
// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (q {{$alias.DownSingular}}Query) UpsertP(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert({{if not .NoContext}}ctx, {{end -}} exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

{{end -}}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q {{$alias.DownSingular}}Query) Upsert(o *{{$alias.UpSingular}}, {{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("{{.PkgName}}: no {{.Table.Name}} provided for upsert")
	}

	{{- template "timestamp_upsert_helper" . }}

	{{if not .NoHooks -}}
	if err := o.doBeforeUpsertHooks({{if not .NoContext}}ctx, {{end -}} exec); err != nil {
		return err
	}
	{{- end}}

	nzDefaults := queries.NonZeroDefaultSet({{$alias.DownSingular}}ColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	{{$alias.DownSingular}}UpsertCacheMut.RLock()
	cache, cached := {{$alias.DownSingular}}UpsertCache[key]
	{{$alias.DownSingular}}UpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			{{$alias.DownSingular}}AllColumns,
			{{$alias.DownSingular}}ColumnsWithDefault,
			{{$alias.DownSingular}}ColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			{{$alias.DownSingular}}AllColumns,
			{{$alias.DownSingular}}PrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("{{.PkgName}}: unable to upsert {{.Table.Name}}, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len({{$alias.DownSingular}}PrimaryKeyColumns))
			copy(conflict, {{$alias.DownSingular}}PrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "{{$schemaTable}}", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping({{$alias.DownSingular}}Type, {{$alias.DownSingular}}Mapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping({{$alias.DownSingular}}Type, {{$alias.DownSingular}}Mapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	{{if .NoContext -}}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	{{else -}}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	{{end -}}

	if len(cache.retMapping) != 0 {
		{{if .NoContext -}}
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		{{else -}}
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		{{end -}}
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		{{if .NoContext -}}
		_, err = exec.Exec(cache.query, vals...)
		{{else -}}
		_, err = exec.ExecContext(ctx, cache.query, vals...)
		{{end -}}
	}
	if err != nil {
		return errors.Wrap(err, "{{.PkgName}}: unable to upsert {{.Table.Name}}")
	}

	if !cached {
		{{$alias.DownSingular}}UpsertCacheMut.Lock()
		{{$alias.DownSingular}}UpsertCache[key] = cache
		{{$alias.DownSingular}}UpsertCacheMut.Unlock()
	}

	{{if not .NoHooks -}}
	return o.doAfterUpsertHooks({{if not .NoContext}}ctx, {{end -}} exec)
	{{- else -}}
	return nil
	{{- end}}
}
