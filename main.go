package main

import (
	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/text/cases"
)

type Field struct {
	Name          string
	Type          string
	IsPrimaryKey  bool
	Nullable      bool
	Default       interface{}
	Index         []Index
	Length        int
	Comment       string
	AutoIncrement bool
}

type Index struct {
	Name     string
	Priority int
	Uniq     bool
}

type Table struct {
	Name   string
	Fields []Field
}

func main() {
	var driver string
	err := singleSelect("Choose your database", []string{"mysql", "postgres", "sqlite"}, &driver)
	if err != nil {
		return
	}

	table := Table{}
	err = fill("Table Name: ", &table.Name)
	if err != nil {
		return
	}

	for {
		f := Field{}
		if err = fill("Field Name: ", &f.Name); err != nil {
			return
		}
		

	}
	

	// switch driver {
	// case "mysql":
	// case "postgres":
	// case "sqlite":
	// default:
	// 	panic("unsupported driver")
	// }
}


func singleSelect(msg string, options []string, ans interface{}) error {
	p := survey.Select{
		Message: msg,
		Options: options,
	}
	return survey.AskOne(&p, ans)
}

func fill(msg string, ans interface{}) error {
	p := &survey.Input{
		Message: msg,
	}
	return survey.AskOne(p, ans)
}

func getStringType(driver string) []string {
	switch driver {
	case "mysql":
		return []string{"CHAR", "VARCHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT"}
	case "postgres":
		return []string{"CHAR", "VARCHAR", "TEXT"}
	default:
		return nil
	}
}

func getNumericType(driver string) []string {
	switch driver {
	case "mysql":
		return []string{"TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT", "FLOAT", "DOUBLE", "DECIMAL"}
	case "postgres":
		return []string{"SMALLINT", "INT", "BIGINT", "DECIMAL", "SMALLSERIAL", "SERIAL", "BIGSERIAL"}
	default:
		return nil
	}
}
