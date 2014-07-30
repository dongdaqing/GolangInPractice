package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	// "reflect"
)

func main() {
	excelFileName := "../../conf/testfile.xlsx"
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		fmt.Println("there is error")

	}
	for _, sheet := range xlFile.Sheets {
		//打印变量类型
		// v := reflect.TypeOf(sheet.Rows)
		// fmt.Println(v)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
			//打印第二列
			// fmt.Println(row.Cells[1])
			// fmt.Println(len(row.Cells))
		}
	}
}
