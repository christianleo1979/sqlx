package sqlx

import "database/sql"

func RowsFromExternal(r *sql.Rows, db *DB) *Rows {
	return &Rows{Rows: r, unsafe: db.unsafe, Mapper: db.Mapper}
}

func RowFromExternal(r *sql.Rows, err error, db *DB) *Row {
	return &Row{rows: r, err: err, unsafe: db.unsafe, Mapper: db.Mapper}
}
