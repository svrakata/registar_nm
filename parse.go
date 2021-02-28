package main

import (
	"fmt"
	"path/filepath"

	"github.com/tealeg/xlsx/v3"
)

var sourceName = "Ek_atte.xlsx"

// ParseEkatte parses the source xlsx file
func ParseEkatte() error {
	dirPath, err := getPaths()

	if err != nil {
		return err
	}

	wb, err := xlsx.OpenFile(filepath.Join(dirPath, sourceName))

	if err != nil {
		return err
	}

	sheetName := wb.Sheets[0].Name
	sheet, ok := wb.Sheet[sheetName]

	if !ok {
		fmt.Println("Sheet doesnt exist")
	}

	defer func() {
		sheet.Close()
	}()

	fmt.Println(sheet)
	return nil
}

// struct

func rowVisitior(r *xlsx.Row) error {

}
