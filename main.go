package main

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	_ "github.com/lib/pq"
	// for pgx, need to use it's standard library compatibility driver:
	// _ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if len(os.Args) <= 1 {
		log.Fatalln("Please provide SQL SELECT statement as first argument")
	}
	stmt := os.Args[1]

	// change driveName to pgx if using pgx
	db, err := sql.Open("postgres", "postgres://root@127.0.0.1:26257/defaultdb?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := db.Query(stmt)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	// get column names
	cols, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return
	}

	// use tabwriter to print result with ease
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)
	defer w.Flush()

	sep := []byte("\t")
	newLine := []byte("\n")

	// print column names
	w.Write([]byte(strings.Join(cols, "\t") + "\n"))

	// row is the storage for row.Scans.
	// for printing, []byte is sufficient for most of the data type,
	// including NULLABLE column, which will be empty
	row := make([][]byte, len(cols))

	// row.Scan accepts variadic POINTER (of any type) arguments.
	// to make it able to take dynamic number of arguments,
	// we need a []any variable and use it in ellipsis operator.
	// rowPtr is just a slice of pointers pointing to every element in row
	rowPtr := make([]any, len(cols))
	for i := range row {
		rowPtr[i] = &row[i]
	}

	for rows.Next() {
		// through rowPtr, we can scan data to row.
		// without rowPtr, we need to do very complex type assertion
		// and dereference when printing (e.g. *(rowPtr[i].(*[]byte)))
		// we also cannot use the helper function (bytes.Join) below
		err := rows.Scan(rowPtr...)
		if err != nil {
			log.Println(err)
			return
		}

		w.Write(bytes.Join(row, sep))
		w.Write(newLine)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return
	}
}
