package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
)

func main() {
	xlsx := excelize.CreateFile()
	xlsx.NewSheet(2, "Sheet2")
	xlsx.NewSheet(3, "Sheet3")
	xlsx.SetCellInt("Sheet2", "A23", 10)
	xlsx.SetCellInt("Sheet2", "A1", 1)
	xlsx.SetCellInt("Sheet2", "A2", 2)
	xlsx.SetCellStr("Sheet3", "B20", "Hello")
	err := xlsx.WriteTo("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
