package ekatte

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/tealeg/xlsx/v3"
)

var sourceName = "Ek_atte.xlsx"

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	return nil
}

func rowVisitor(r *xlsx.Row) error {
	r.ForEachCell(cellVisitor)

	return nil
}

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
	sh, ok := wb.Sheet[sheetName]

	if !ok {
		fmt.Println("Sheet doesnt exist")
	}

	sh.ForEachRow(rowVisitor)

	defer func() {
		sh.Close()
	}()

	return nil
}

// struct
