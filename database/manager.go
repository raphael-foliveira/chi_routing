package database

import (
	"database/sql"
	"errors"
	"reflect"
)

type manager struct {
	target interface{}
}

func Manager(target interface{}) *manager {
	return &manager{
		target: target,
	}
}

func (m *manager) QueryRows(query string, params ...any) {
	rows, err := Db.Query(query, params...)
	if err != nil {
		panic(err)
	}
	m.ScanRows(rows)
}

func (m *manager) QueryRow(query string, params ...any) {
	row := Db.QueryRow(query, params...)
	m.ScanRow(row)
}

func (m *manager) ScanRows(rows *sql.Rows) error {
	if reflect.TypeOf(m.target).Kind() != reflect.Ptr || reflect.TypeOf(m.target).Elem().Kind() != reflect.Slice {
		return errors.New("Target must be a pointer to a slice")
	}

	sliceVal := reflect.ValueOf(m.target).Elem() // []students.Student{} (value)
	structType := sliceVal.Type().Elem()         // students.Student (type)

	for rows.Next() {
		structVal := reflect.New(structType).Elem() // new students.Student{} (value)
		numFields := structVal.NumField()           // number of fields in Student struct (int)
		scanVals := make([]interface{}, numFields)  //interface to receive pointers

		for i := 0; i < numFields; i++ {
			// .Addr() works because `result` is a reference (&) when we call reflect.ValueOf(result)
			scanVals[i] = structVal.Field(i).Addr().Interface()
		}

		if err := rows.Scan(scanVals...); err != nil {
			return err
		}

		sliceVal.Set(reflect.Append(sliceVal, structVal)) // put rows values in `result`
	}
	return nil
}

func (m *manager) ScanRow(row *sql.Row) error {
	if reflect.TypeOf(m.target).Kind() != reflect.Ptr || reflect.TypeOf(m.target).Elem().Kind() != reflect.Struct {
		return errors.New("Target must be a pointer to a struct")
	}

	structType := reflect.TypeOf(m.target).Elem()
	structVal := reflect.New(structType).Elem()
	numFields := structVal.NumField()
	scanVals := make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		scanVals[i] = structVal.Field(i).Addr().Interface()
	}
	err := row.Scan(scanVals...)
	if err != nil {
		return err
	}
	reflect.ValueOf(m.target).Elem().Set(structVal) // put row values in `result`
	return nil
}
