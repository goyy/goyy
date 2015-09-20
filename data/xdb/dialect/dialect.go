// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

type Table struct {
	Id      string
	Name    string
	Comment string
}

type Column struct {
	Id       string
	Name     string
	Comment  string
	Nullable bool
	Default  string
	Key      string
	Domain
}

type Domain struct {
	Id        string
	Name      string
	Type      string
	Length    int
	Precision int
}

// Index represents a table index and is returned via the Indexed interface.
type Index struct {
	Name    string
	Columns []string
	Unique  bool
}

type Dialect interface {

	// Quote will quote identifiers in a SQL statement.
	Quote(s string) string

	// SqlType returns the SQL type for the provided interface type.
	SqlType(domain Domain) string

	// CreateTable returns the sql for creating a table.
	CreateTable(table *Table, columns ...*Column) string

	// DropTable returns the sql for dropping the specified table.
	DropTable(table string) string

	// RenameTable returns the sql for renaming the specified table.
	RenameTable(from, to string) string

	// AddColumn returns the sql for adding the specified column in table.
	AddColumn(table string, column *Column) string

	// RenameColumn returns the sql for renaming the specified column in table.
	RenameColumn(table, from, to string) string

	// ChangeColumn returns the sql for changing the column data type.
	ChangeColumn(table string, column *Column) string

	// DropColumn returns the sql for removing the column.
	DropColumn(table, column string) string

	// CreateIndex returns the sql for creating an index on the specified column.
	CreateIndex(name, table string, unique bool, columns ...string) string

	// DropIndex returns the sql for dropping the index.
	DropIndex(name string) string
}
